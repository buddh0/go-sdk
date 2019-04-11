package rpc

import "github.com/binance-chain/go-sdk/common/types"

func (c *HTTP) GetStakeValidators() ([]types.Validator, error) {
	rawVal, err := c.ABCIQuery("custom/stake/validators", nil)
	if err != nil {
		return nil, err
	}
	var validators []types.Validator
	err = c.cdc.UnmarshalJSON(rawVal.Response.GetValue(), &validators)
	return validators, err

}

func (c *HTTP) GetDelegatorUnbondingDelegations(delegatorAddr types.AccAddress) ([]types.UnbondingDelegation, error) {
	param := struct {
		DelegatorAddr types.AccAddress
	}{delegatorAddr}
	bz, err := c.cdc.MarshalJSON(param)
	if err != nil {
		return nil, err
	}

	rawDel, err := c.ABCIQuery("custom/stake/delegatorUnbondingDelegations", bz)
	if err != nil {
		return nil, err
	}
	var unbondingDelegations []types.UnbondingDelegation
	err = c.cdc.UnmarshalJSON(rawDel.Response.GetValue(), &unbondingDelegations)
	return unbondingDelegations, err

}
