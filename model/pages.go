package model

type Pages struct {
	UserNames []string
	Reports []*Report
	PageNo  int64//当前页
	PageSize int64//每页的条数
	TotalPageNo int64//总页数
	TotalRecord int64//总记录数
	Time string
	UserName string
}
func (p *Pages)GetNextPage()int64{
	if p.PageNo<p.TotalPageNo{
		return p.PageNo+1
	}else{
		return p.TotalPageNo
	}
}
func (p *Pages)GetPrevPage()int64{
	if p.PageNo==1{
		return 1
	}else{
		return p.PageNo-1
	}
}
