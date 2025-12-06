package cmd

import (
	"os"
	"time"

	"dns-scanner/pkg/output"
	"dns-scanner/pkg/scanner"
	"dns-scanner/pkg/wordlist"

	"github.com/spf13/cobra"
)

var (
	domain       string
	wordlistPath string
	outputPath   string
	format       string
	concurrency  int
	dnsServer    string
	timeout      int
	enableIPv6   bool
)

var rootCmd = &cobra.Command{
	Use:   "dns-scanner",
	Short: "DNS 子域名枚举扫描器",
	Long: `DNS Scanner 是一个用于发现目标域名子域名的工具。
通过字典爆破的方式枚举可能存在的子域名。

示例:
  dns-scanner -d example.com
  dns-scanner -d example.com -w wordlist.txt -o result.json -f json`,
	Run: runScan,
}

func init() {
	rootCmd.Flags().StringVarP(&domain, "domain", "d", "", "目标域名 (必填)")
	rootCmd.Flags().StringVarP(&wordlistPath, "wordlist", "w", "", "自定义字典文件路径")
	rootCmd.Flags().StringVarP(&outputPath, "output", "o", "", "输出文件路径")
	rootCmd.Flags().StringVarP(&format, "format", "f", "txt", "输出格式 (json/csv/txt)")
	rootCmd.Flags().IntVarP(&concurrency, "concurrency", "c", 50, "并发数")
	rootCmd.Flags().StringVar(&dnsServer, "dns", "8.8.8.8:53", "DNS 服务器地址")
	rootCmd.Flags().IntVarP(&timeout, "timeout", "t", 3, "DNS 查询超时时间 (秒)")
	rootCmd.Flags().BoolVarP(&enableIPv6, "ipv6", "6", false, "同时扫描 IPv6 (AAAA 记录)")

	rootCmd.MarkFlagRequired("domain")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func runScan(cmd *cobra.Command, args []string) {
	// 加载字典
	words, err := wordlist.Load(wordlistPath)
	if err != nil {
		output.PrintError("加载字典失败: " + err.Error())
		os.Exit(1)
	}

	output.PrintStart(domain, len(words))

	// 创建扫描器
	s := scanner.New(
		domain,
		dnsServer,
		time.Duration(timeout)*time.Second,
		concurrency,
	)
	s.EnableIPv6 = enableIPv6

	// 设置发现回调
	s.OnFound = func(result scanner.Result) {
		output.PrintResult(result)
	}

	// 执行扫描
	results := s.Scan(words)

	// 打印汇总
	output.PrintSummary(domain, len(words), len(results))

	// 输出到文件
	if outputPath != "" {
		writer := output.New(format, outputPath)
		if err := writer.Write(domain, results); err != nil {
			output.PrintError("写入文件失败: " + err.Error())
			os.Exit(1)
		}
		output.PrintInfo("结果已保存到: " + outputPath)
	}
}
