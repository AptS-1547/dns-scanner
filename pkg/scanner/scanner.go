package scanner

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/miekg/dns"
)

// Result 扫描结果
type Result struct {
	Subdomain string   `json:"subdomain"`
	IPs       []string `json:"ips"`
	IPv6      []string `json:"ipv6,omitempty"`
}

// Scanner DNS 扫描器
type Scanner struct {
	Domain      string
	DNSServer   string
	Timeout     time.Duration
	Concurrency int
	EnableIPv6  bool           // 是否扫描 IPv6
	OnFound     func(Result)   // 发现子域名时的回调
}

// New 创建新的扫描器
func New(domain, dnsServer string, timeout time.Duration, concurrency int) *Scanner {
	return &Scanner{
		Domain:      domain,
		DNSServer:   dnsServer,
		Timeout:     timeout,
		Concurrency: concurrency,
	}
}

// Scan 执行扫描
func (s *Scanner) Scan(wordlist []string) []Result {
	var results []Result
	var mu sync.Mutex

	jobs := make(chan string, s.Concurrency)
	var wg sync.WaitGroup

	// 启动 worker
	for i := 0; i < s.Concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for word := range jobs {
				subdomain := fmt.Sprintf("%s.%s", word, s.Domain)
				ipv4s, ipv6s := s.resolve(subdomain)
				if len(ipv4s) > 0 || len(ipv6s) > 0 {
					result := Result{
						Subdomain: subdomain,
						IPs:       ipv4s,
						IPv6:      ipv6s,
					}
					mu.Lock()
					results = append(results, result)
					mu.Unlock()

					if s.OnFound != nil {
						s.OnFound(result)
					}
				}
			}
		}()
	}

	// 分发任务
	for _, word := range wordlist {
		jobs <- word
	}
	close(jobs)

	wg.Wait()
	return results
}

// resolve 解析域名获取 IP
func (s *Scanner) resolve(domain string) ([]string, []string) {
	c := new(dns.Client)
	c.Timeout = s.Timeout

	var ipv4s, ipv6s []string

	// 查询 A 记录 (IPv4)
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(domain), dns.TypeA)
	m.RecursionDesired = true

	r, _, err := c.Exchange(m, s.DNSServer)
	if err == nil && r.Rcode == dns.RcodeSuccess {
		for _, ans := range r.Answer {
			if a, ok := ans.(*dns.A); ok {
				ipv4s = append(ipv4s, a.A.String())
			}
		}
	}

	// 查询 AAAA 记录 (IPv6)
	if s.EnableIPv6 {
		m6 := new(dns.Msg)
		m6.SetQuestion(dns.Fqdn(domain), dns.TypeAAAA)
		m6.RecursionDesired = true

		r6, _, err := c.Exchange(m6, s.DNSServer)
		if err == nil && r6.Rcode == dns.RcodeSuccess {
			for _, ans := range r6.Answer {
				if aaaa, ok := ans.(*dns.AAAA); ok {
					ipv6s = append(ipv6s, aaaa.AAAA.String())
				}
			}
		}
	}

	return ipv4s, ipv6s
}

// ScanWithProgress 带进度的扫描
func (s *Scanner) ScanWithProgress(wordlist []string, onProgress func(current, total int)) []Result {
	var results []Result
	var mu sync.Mutex
	var progressMu sync.Mutex
	progress := 0
	total := len(wordlist)

	jobs := make(chan string, s.Concurrency)
	var wg sync.WaitGroup

	// 启动 worker
	for i := 0; i < s.Concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for word := range jobs {
				subdomain := fmt.Sprintf("%s.%s", word, s.Domain)
				ipv4s, ipv6s := s.resolve(subdomain)

				progressMu.Lock()
				progress++
				currentProgress := progress
				progressMu.Unlock()

				if onProgress != nil {
					onProgress(currentProgress, total)
				}

				if len(ipv4s) > 0 || len(ipv6s) > 0 {
					result := Result{
						Subdomain: subdomain,
						IPs:       ipv4s,
						IPv6:      ipv6s,
					}
					mu.Lock()
					results = append(results, result)
					mu.Unlock()

					if s.OnFound != nil {
						s.OnFound(result)
					}
				}
			}
		}()
	}

	// 分发任务
	for _, word := range wordlist {
		jobs <- word
	}
	close(jobs)

	wg.Wait()
	return results
}

// FormatIPs 格式化 IP 列表
func FormatIPs(ips []string) string {
	return strings.Join(ips, ", ")
}
