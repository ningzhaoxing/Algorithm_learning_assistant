package models

// Problem 存储力扣题目信息

// LeetCodeProfile 存储力扣用户信息
type LeetCodeProfile struct {
	Username    string    // 用户名
	SolvedCount string    // 解题数量
	Problems    []Problem // 解决的题目列表
}

func NewLeetCodeProfile(username, solvedCount string, problems []Problem) *LeetCodeProfile {
	name := map[string]string{
		"kan-fan-xing":             "宁赵星",
		"xun_xun":                  "蒋睿勋",
		"zao-an-e":                 "雪怡琦",
		"gui-tu-960":               "方腾飞",
		"ding-mao-s":               "田家杰",
		"trusting-6rothendieckqgx": "贺丽帆",
		"6oofy-gangulyxsi":         "韩硕博",
		"practical-snyderqvy":      "王怡晗",
		"hardcore-swirlesrz0":      "王玉龙",
		"festive-i2ubinwnk":        "李壮",
	}
	return &LeetCodeProfile{
		Username:    name[username],
		SolvedCount: solvedCount,
		Problems:    problems,
	}
}
