package poi

import (
	"github.com/micro-plat/wechat/mp"
)

// Delete 删除门店.
func Delete(clt *mp.Context, poiId int64) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/poi/delpoi?access_token="

	var request = struct {
		PoiId int64 `json:"poi_id"`
	}{
		PoiId: poiId,
	}
	var result mp.Error
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}
