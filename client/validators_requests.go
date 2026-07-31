package client

type ValidatorsQuery struct {
	Pagination
	Address                string  `url:"address,omitempty"`
	Attributes             string  `url:"attributes,omitempty"`
	BlocksLastHash         string  `url:"blocks.last.hash,omitempty"`
	BlocksLastNumber       int64   `url:"blocks.last.number,omitempty"`
	BlocksProduced         int64   `url:"blocks.produced,omitempty"`
	ForgedFees             string  `url:"forged.fees,omitempty"`
	ForgedFeesFrom         string  `url:"forged.fees.from,omitempty"`
	ForgedFeesTo           string  `url:"forged.fees.to,omitempty"`
	ForgedRewards          string  `url:"forged.rewards,omitempty"`
	ForgedRewardsFrom      string  `url:"forged.rewards.from,omitempty"`
	ForgedRewardsTo        string  `url:"forged.rewards.to,omitempty"`
	ForgedTotal            string  `url:"forged.total,omitempty"`
	ForgedTotalFrom        string  `url:"forged.total.from,omitempty"`
	ForgedTotalTo          string  `url:"forged.total.to,omitempty"`
	IsResigned             bool    `url:"isResigned,omitempty"`
	ProductionApproval     float64 `url:"production.approval,omitempty"`
	ProductionApprovalFrom float64 `url:"production.approval.from,omitempty"`
	ProductionApprovalTo   float64 `url:"production.approval.to,omitempty"`
	PublicKey              string  `url:"publicKey,omitempty"`
	Rank                   int64   `url:"rank,omitempty"`
	RankFrom               int64   `url:"rank.from,omitempty"`
	RankTo                 int64   `url:"rank.to,omitempty"`
	Username               string  `url:"username,omitempty"`
	Votes                  string  `url:"votes,omitempty"`
	VotesFrom              string  `url:"votes.from,omitempty"`
	VotesTo                string  `url:"votes.to,omitempty"`
	OrderBy                string  `url:"orderBy,omitempty"`
}
