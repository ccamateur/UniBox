package main_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"main.go/swap"
)

func TestSwapFunctions(t *testing.T) {

	walletAddress := ""
	privateKeyHex := ""

	//intialize web3client
	swap.InitializeWeb3Client("https://eth-ropsten.alchemyapi.io/v2/OKhklavMx34YUKkZEIwsS9bMgw0EBJnV")
	//initialize uniswapRouter02 with wallet address
	swap.InitializeUniswapRouter02(walletAddress)

	//test new wallet key for signing
	key := swap.NewWalletKey(privateKeyHex)
	if key == nil {
		t.Fatal("NewWalletKey Error")
	}

	//set lp path
	path := []common.Address{common.HexToAddress("0x000000000"), common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48")}
	//reciever wallet address
	to := common.HexToAddress(walletAddress)
	//exact amount out
	// amountOut := uint(100)
	//minimum amount out
	amountOutMin := uint(0)
	//exact amount in
	// amountIn := uint(100)
	//maximum amount in
	// amountInMax := uint(200)
	//timeout deadline
	deadline := uint(1000000)

	// //SwapExactTokensForTokens
	// txn0 := swap.SwapExactTokensForTokens(amountIn, amountOutMin, path, to, deadline)
	// if txn0 == nil {
	// 	t.Fatal("SwapExactTokensForTokens Error")
	// }
	// // txn0.SignAndSend(key, 3)

	// //SwapTokensForExactTokens
	// txn1 := swap.SwapTokensForExactTokens(amountOut, amountInMax, path, to, deadline)
	// if txn1 == nil {
	// 	t.Fatal("SwapTokensForExactTokens Error")
	// }

	//SwapExactETHForTokens
	txn2 := swap.SwapExactETHForTokens(amountOutMin, path, to, deadline)
	if txn2 == nil {
		t.Fatal("SwapExactETHForTokens Error")
	}
	err := txn2.SignAndSend(key, 3)
	if err != nil {
		t.Fatal("Sign and Send Error", err)
	}

	// //SwapTokensForExactETH
	// txn3 := swap.SwapTokensForExactETH(amountOut, amountInMax, path, to, deadline)
	// if txn3 == nil {
	// 	t.Fatal("SwapTokensForExactETH Error")
	// }

	// //SwapExactTokensForETH
	// txn4 := swap.SwapExactTokensForETH(amountIn, amountOutMin, path, to, deadline)
	// if txn4 == nil {
	// 	t.Fatal("SwapExactTokensForETH Error")
	// }

	// //SwapETHForExactTokens
	// txn5 := swap.SwapETHForExactTokens(amountOut, path, to, deadline)
	// if txn5 == nil {
	// 	t.Fatal("SwapETHForExactTokens Error")
	// }

}

// func TestSwapFunctionsWithFeeOnTransfer(t *testing.T) {
// 	//test new wallet key for signing
// 	key := swap.NewWalletKey("")
// 	if key == nil {
// 		t.Fatal("NewWalletKey Error")
// 	}

// 	//set lp path
// 	path := []common.Address{common.HexToAddress(""), common.HexToAddress("")}
// 	//reciever wallet address
// 	to := common.HexToAddress(walletAddress)
// 	//minimum amount out
// 	amountOutMin := uint(0)
// 	//exact amount in
// 	amountIn := uint(100)
// 	//timeout deadline
// 	deadline := uint(1000)

// 	//SwapExactTokensForTokensSupportingFeeOnTransferTokens
// 	txn6 := swap.SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn, amountOutMin, path, to, deadline)
// 	if txn6 == nil {
// 		t.Fatal("SwapExactTokensForTokensSupportingFeeOnTransferTokens Error")
// 	}

// 	//SwapExactETHForTokensSupportingFeeOnTransferTokens
// 	txn7 := swap.SwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin, path, to, deadline)
// 	if txn7 == nil {
// 		t.Fatal("SwapExactETHForTokensSupportingFeeOnTransferTokens Error")
// 	}

// 	//SwapExactTokensForETHSupportingFeeOnTransferTokens
// 	txn8 := swap.SwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn, amountOutMin, path, to, deadline)
// 	if txn8 == nil {
// 		t.Fatal("SwapExactTokensForETHSupportingFeeOnTransferTokens Error")
// 	}

// }
