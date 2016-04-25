package main

type lineReport struct {
	line  string
	count int
	files []string
}

type lineReports []lineReport

func (rs lineReports) Len() int {
	return len(rs)
}

func (rs lineReports) Less(i, j int) bool {
	if rs[i].count > rs[j].count {
		return true
	} else if rs[i].count < rs[j].count {
		return false
	} else {
		return rs[i].line > rs[j].line
	}
}

func (rs lineReports) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}
