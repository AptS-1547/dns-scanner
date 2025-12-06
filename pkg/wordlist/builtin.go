package wordlist

// BuiltinWordlist 内置常用子域名字典
var BuiltinWordlist = []string{
	// 通用
	"www", "mail", "ftp", "localhost", "webmail", "smtp", "pop", "ns1", "ns2",
	"dns", "dns1", "dns2", "mx", "mx1", "mx2", "imap", "pop3", "admin", "secure",
	"vpn", "api", "dev", "staging", "test", "beta", "demo", "app", "apps",
	"web", "portal", "login", "m", "mobile", "wap", "static", "assets", "cdn",
	"img", "images", "image", "css", "js", "media", "download", "downloads",
	"upload", "uploads", "files", "file", "doc", "docs", "support", "help",
	"status", "blog", "news", "shop", "store", "cart", "checkout", "pay",
	"payment", "billing", "invoice", "order", "orders", "account", "accounts",
	"my", "user", "users", "member", "members", "client", "clients", "customer",
	"customers", "partner", "partners", "affiliate", "affiliates", "reseller",

	// 管理后台
	"admin", "administrator", "adm", "manage", "manager", "management", "cms",
	"backend", "backoffice", "dashboard", "control", "cpanel", "panel", "wp-admin",
	"phpmyadmin", "pma", "mysql", "db", "database", "sql", "mssql", "postgres",
	"oracle", "mongo", "mongodb", "redis", "elastic", "elasticsearch", "kibana",

	// 开发测试
	"dev", "devel", "develop", "development", "test", "testing", "qa", "uat",
	"staging", "stage", "stg", "demo", "sandbox", "beta", "alpha", "preview",
	"pre", "preprod", "pre-prod", "prod", "production", "live", "release",
	"local", "localhost", "debug", "ci", "cd", "jenkins", "gitlab", "github",
	"git", "svn", "repo", "repository", "build", "deploy", "docker", "k8s",
	"kubernetes", "rancher", "traefik", "nginx", "apache", "tomcat", "jboss",

	// 邮件相关
	"mail", "email", "e-mail", "webmail", "smtp", "pop", "pop3", "imap",
	"mx", "mx1", "mx2", "mx3", "mailserver", "mail1", "mail2", "mail3",
	"relay", "mailrelay", "mailgateway", "exchange", "outlook", "postfix",
	"sendmail", "mailman", "list", "lists", "mailing", "newsletter",

	// 网络设备
	"router", "gateway", "firewall", "switch", "proxy", "lb", "loadbalancer",
	"load-balancer", "haproxy", "f5", "citrix", "vpn", "openvpn", "ipsec",
	"ssl", "sslvpn", "remote", "rdp", "ssh", "sftp", "ftp", "tftp",

	// DNS相关
	"ns", "ns1", "ns2", "ns3", "ns4", "dns", "dns1", "dns2", "dns3",
	"nameserver", "nameserver1", "nameserver2", "resolver", "pdns",

	// 云服务
	"cloud", "aws", "azure", "gcp", "google", "alibaba", "aliyun", "tencent",
	"qcloud", "huawei", "hwcloud", "s3", "oss", "cos", "blob", "storage",
	"bucket", "cdn", "edge", "cloudfront", "akamai", "fastly", "cloudflare",

	// 监控安全
	"monitor", "monitoring", "nagios", "zabbix", "prometheus", "grafana",
	"log", "logs", "logging", "elk", "splunk", "siem", "soc", "security",
	"waf", "ids", "ips", "antivirus", "av", "scan", "scanner", "audit",

	// 服务与应用
	"api", "api1", "api2", "rest", "graphql", "soap", "rpc", "grpc",
	"ws", "websocket", "socket", "gateway", "service", "services", "svc",
	"microservice", "microservices", "auth", "oauth", "sso", "ldap", "ad",
	"cas", "saml", "oidc", "jwt", "token", "key", "keys", "cert", "certs",
	"certificate", "ca", "pki", "vault", "secret", "secrets", "config",

	// 内部系统
	"intranet", "internal", "corp", "corporate", "office", "hr", "erp",
	"crm", "oa", "wiki", "confluence", "jira", "gitlab", "github", "bitbucket",
	"slack", "teams", "zoom", "meet", "meeting", "calendar", "schedule",

	// 地区分站
	"cn", "us", "uk", "eu", "asia", "na", "sa", "au", "jp", "kr", "de", "fr",
	"ru", "br", "in", "sg", "hk", "tw", "global", "intl", "international",

	// 数字编号
	"1", "2", "3", "4", "5", "01", "02", "03", "04", "05",
	"server1", "server2", "server3", "host1", "host2", "host3",
	"node1", "node2", "node3", "web1", "web2", "web3",
	"app1", "app2", "app3", "db1", "db2", "db3",

	// 其他常见
	"home", "main", "default", "index", "root", "public", "private",
	"old", "new", "backup", "bak", "archive", "temp", "tmp", "cache",
	"proxy", "reverse", "forward", "origin", "upstream", "downstream",
	"primary", "secondary", "master", "slave", "replica", "standby",
	"active", "passive", "hot", "cold", "warm", "dr", "disaster", "recovery",
}
