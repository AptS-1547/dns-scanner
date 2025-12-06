package output

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"dns-scanner/pkg/scanner"
)

// ScanOutput 扫描输出结构
type ScanOutput struct {
	Domain string           `json:"domain"`
	Total  int              `json:"total"`
	Found  []scanner.Result `json:"found"`
}

// Writer 输出写入器
type Writer struct {
	Format   string
	FilePath string
}

// New 创建输出写入器
func New(format, filepath string) *Writer {
	return &Writer{
		Format:   format,
		FilePath: filepath,
	}
}

// Write 写入结果
func (w *Writer) Write(domain string, results []scanner.Result) error {
	output := ScanOutput{
		Domain: domain,
		Total:  len(results),
		Found:  results,
	}

	switch strings.ToLower(w.Format) {
	case "json":
		return w.writeJSON(output)
	case "csv":
		return w.writeCSV(output)
	case "txt":
		return w.writeTXT(output)
	default:
		return w.writeTXT(output)
	}
}

// writeJSON 写入 JSON 格式
func (w *Writer) writeJSON(output ScanOutput) error {
	data, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(w.FilePath, data, 0644)
}

// writeCSV 写入 CSV 格式
func (w *Writer) writeCSV(output ScanOutput) error {
	file, err := os.Create(w.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// 写入表头
	if err := writer.Write([]string{"Subdomain", "IPv4", "IPv6"}); err != nil {
		return err
	}

	// 写入数据
	for _, result := range output.Found {
		ipv4 := strings.Join(result.IPs, ";")
		ipv6 := strings.Join(result.IPv6, ";")
		if err := writer.Write([]string{result.Subdomain, ipv4, ipv6}); err != nil {
			return err
		}
	}

	return nil
}

// writeTXT 写入 TXT 格式
func (w *Writer) writeTXT(output ScanOutput) error {
	file, err := os.Create(w.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintf(file, "# DNS Scan Results for %s\n", output.Domain)
	fmt.Fprintf(file, "# Total found: %d\n\n", output.Total)

	for _, result := range output.Found {
		ipv4 := strings.Join(result.IPs, ", ")
		fmt.Fprintf(file, "%s -> %s", result.Subdomain, ipv4)
		if len(result.IPv6) > 0 {
			ipv6 := strings.Join(result.IPv6, ", ")
			fmt.Fprintf(file, " | IPv6: %s", ipv6)
		}
		fmt.Fprintf(file, "\n")
	}

	return nil
}

// PrintResult 打印单个结果到终端
func PrintResult(result scanner.Result) {
	ipv4 := strings.Join(result.IPs, ", ")
	if len(result.IPv6) > 0 {
		ipv6 := strings.Join(result.IPv6, ", ")
		fmt.Printf("\033[32m[+]\033[0m %s -> %s \033[35m| IPv6: %s\033[0m\n", result.Subdomain, ipv4, ipv6)
	} else {
		fmt.Printf("\033[32m[+]\033[0m %s -> %s\n", result.Subdomain, ipv4)
	}
}

// PrintSummary 打印汇总信息
func PrintSummary(domain string, total, found int) {
	fmt.Printf("\n\033[36m[*]\033[0m Scan completed for %s\n", domain)
	fmt.Printf("\033[36m[*]\033[0m Checked: %d | Found: %d\n", total, found)
}

// PrintStart 打印开始信息
func PrintStart(domain string, wordlistCount int) {
	fmt.Printf("\033[36m[*]\033[0m Scanning %s with %d subdomains...\n", domain, wordlistCount)
}

// PrintError 打印错误信息
func PrintError(msg string) {
	fmt.Printf("\033[31m[!]\033[0m %s\n", msg)
}

// PrintInfo 打印普通信息
func PrintInfo(msg string) {
	fmt.Printf("\033[36m[*]\033[0m %s\n", msg)
}
