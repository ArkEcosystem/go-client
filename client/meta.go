package client

type Meta struct {
	Count                uint32  `url:"count,omitempty"`
	PageCount            uint32  `url:"pageCount,omitempty"`
	TotalCount           uint32  `url:"totalCount,omitempty"`
	Next                 *string `url:"next,omitempty"`
	Previous             *string `url:"previous,omitempty"`
	Self                 string  `url:"self,omitempty"`
	First                string  `url:"first,omitempty"`
	Last                 string  `url:"last,omitempty"`
	TotalCountIsEstimate bool    `url:"totalCountIsEstimate,omitempty"`
}
