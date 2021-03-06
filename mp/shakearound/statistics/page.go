package statistics

import (
	"github.com/micro-plat/wechat/mp"
)

// 以页面为维度的数据统计接口
func Page(clt *mp.Context, pageId, beginDate, endDate int64) (data []StatisticsBase, err error) {
	request := struct {
		PageId    int64 `json:"page_id"`
		BeginDate int64 `json:"begin_date"`
		EndDate   int64 `json:"end_date"`
	}{
		PageId:    pageId,
		BeginDate: beginDate,
		EndDate:   endDate,
	}

	var result struct {
		mp.Error
		Data []StatisticsBase `json:"data"`
	}

	incompleteURL := "https://api.weixin.qq.com/shakearound/statistics/page?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	data = result.Data
	return
}
