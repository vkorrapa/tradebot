package service

import (
	"github.com/jeremyhahn/tradebot/common"
	"github.com/jeremyhahn/tradebot/dao"
	"github.com/jeremyhahn/tradebot/mapper"
)

type EthereumServiceImpl struct {
	ctx             common.Context
	userDAO         dao.UserDAO
	userMapper      mapper.UserMapper
	ethereumService EthereumService
	EthereumService
}

func NewEthereumService(ctx common.Context, userDAO dao.UserDAO, userMapper mapper.UserMapper,
	marketcapService MarketCapService) (EthereumService, error) {

	var service *EthereumServiceImpl
	if ctx.GetEthereumMode() == "native" {
		gethService, err := NewGethService(ctx, dao.NewUserDAO(ctx), mapper.NewUserMapper())
		if err != nil {
			return nil, err
		}
		service = &EthereumServiceImpl{
			ctx:             ctx,
			userDAO:         userDAO,
			userMapper:      userMapper,
			ethereumService: gethService.(EthereumService)}
		//} else if ctx.GetEthereumMode() == "etherscan" {
	} else {
		etherscanService, err := NewEtherscanService(ctx, userDAO, userMapper, marketcapService)
		if err != nil {
			return nil, err
		}
		service = &EthereumServiceImpl{
			ctx:             ctx,
			userDAO:         userDAO,
			userMapper:      userMapper,
			ethereumService: etherscanService}
	}
	return service, nil
}

func (facade *EthereumServiceImpl) GetAccounts() ([]common.UserContext, error) {
	return facade.ethereumService.GetAccounts()
}

func (facade *EthereumServiceImpl) Login(username, password string) (common.UserContext, error) {
	return facade.ethereumService.Login(username, password)
}

func (facade *EthereumServiceImpl) Register(username, password string) error {
	return facade.ethereumService.Register(username, password)
}

func (facade *EthereumServiceImpl) GetWallet(address string) (common.UserCryptoWallet, error) {
	return facade.ethereumService.GetWallet(address)
}

func (facade *EthereumServiceImpl) GetTransactions(address string) ([]common.Transaction, error) {
	return facade.ethereumService.GetTransactions(address)
}

func (facade *EthereumServiceImpl) GetToken(walletAddress, contractAddress string) (common.EthereumToken, error) {
	return facade.ethereumService.GetToken(walletAddress, contractAddress)
}

func (facade *EthereumServiceImpl) GetTokenTransactions(contractAddress string) ([]common.Transaction, error) {
	return facade.ethereumService.GetTokenTransactions(contractAddress)
}
