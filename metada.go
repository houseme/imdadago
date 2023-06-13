/*
 *  Copyright ImDaDa-Go Author(https://houseme.github.io/imdada-go/). All Rights Reserved.
 *
 *  This Source Code Form is subject to the terms of the Apache-2.0 License.
 *  If a copy of the MIT was not distributed with this file,
 *  You can obtain one at https://github.com/houseme/imdada-go.
 */

package dada

import (
	"context"

	"github.com/bytedance/sonic"

	"github.com/houseme/imdada-go/domain"
)

// QueryCity query city list
// 获取城市信息列表 http://newopen.imdada.cn/#/development/file/cityList
func (c *Client) QueryCity(ctx context.Context, req *domain.CityListQueryRequest) (resp *domain.CityListQueryResponse, err error) {
	if req == nil {
		c.request.Body = ""
	} else {
		if c.request.Body, err = sonic.MarshalString(req); err != nil {
			return nil, err
		}
	}
	c.log.CtxDebugf(ctx, "QueryCity request data: %s ", c.request.Body)
	if err = c.doRequest(ctx, cityCodeList); err != nil {
		return nil, err
	}
	c.log.CtxDebugf(ctx, "QueryCity response data: %s", string(c.response.Body()))
	if err = sonic.Unmarshal(c.response.Body(), &resp); err != nil {
		return nil, err
	}
	return
}
