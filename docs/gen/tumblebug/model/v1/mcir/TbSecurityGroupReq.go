package mcir

type TbSecurityGroupReq struct {
	ConnectionName	string	`json:"connectionName"`
	Description	string	`json:"description"`
	FirewallRules	[]SpiderSecurityRuleInfo	`json:"firewallRules"`
	Name	string	`json:"name"`
	VNetId	string	`json:"vNetId"`
}