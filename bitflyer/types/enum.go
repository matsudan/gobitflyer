package types

type ProductCode string

type MarketType string

const (
	// MarketTypeSpot represents the market type is spot.
	MarketTypeSpot MarketType = "Spot"

	// MarketTypeFx represents the market type is FX.
	MarketTypeFx MarketType = "FX"

	// MarketTypeFutures represents the market type is Futures
	MarketTypeFutures MarketType = "Futures"
)

type ExchangeHealth string

const (
	// The exchange is operating.
	ExchangeHealthNormal ExchangeHealth = "NORMAL"

	// The exchange is experiencing high traffic.
	ExchangeHealthBusy ExchangeHealth = "BUSY"

	// The exchange is experiencing very heavy traffic.
	ExchangeHealthVeryBusy ExchangeHealth = "VERY BUSY"

	// The exchange is experiencing extremely heavy traffic. There is a possibility that orders will fail or be processed after a delay.
	ExchangeHealthSuperBusy ExchangeHealth = "SUPER BUSY"

	// Orders can not be received.
	ExchangeHealthNoOrder ExchangeHealth = "NO ORDER"

	// The exchange has been stopped. Orders will not be accepted.
	ExchangeHealthStop ExchangeHealth = "STOP"
)

type BoardState string

const (
	// Operating
	BoardStateRunning BoardState = "RUNNING"

	// Suspending
	BoardStateClosed BoardState = "CLOSED"

	// Restarting
	BoardStateStarting BoardState = "STARTING"

	// Performing Itayose
	BoardStatePreOpen BoardState = "PREOPEN"

	// Circuit breaker triggered
	BoardStateCircuitBreak BoardState = "CIRCUIT BREAK"

	// Calculating SQ (special quotation) for Lightning Futures after trades complete
	BoardStateAwaitingSQ BoardState = "AWAITING SQ"

	// Lightning Futures maturity reached
	BoardStateMatured BoardState = "MATURED"
)

type ExchangeStatus string

const (
	// The exchange is operating.
	ExchangeStatusNormal ExchangeStatus = "NORMAL"

	// The exchange is experiencing high traffic.
	ExchangeStatusBusy ExchangeStatus = "BUSY"

	// The exchange is experiencing very heavy traffic.
	ExchangeStatusVeryBusy ExchangeStatus = "VERY BUSY"

	// The exchange is experiencing extremely heavy traffic. There is a possibility that orders will fail or be processed after a delay.
	ExchangeStatusSuperBusy ExchangeStatus = "SUPER BUSY"

	// The exchange has been stopped. Orders will not be accepted.
	ExchangeStatusStop ExchangeStatus = "STOP"
)

type DepositStatus string

const (
	// DepositStatusPending represents that the cash deposit is being processed.
	DepositStatusPending DepositStatus = "PENDING"

	// DepositStatusCompleted represents that the deposit has been completed.
	DepositStatusCompleted DepositStatus = "COMPLETED"
)

type WithdrawalStatus string

const (
	// WithdrawalStatusPending represents that the withdrawal is being processed.
	WithdrawalStatusPending WithdrawalStatus = "PENDING"

	// WithdrawalStatusCompleted represents that the withdrawal has been completed.
	WithdrawalStatusCompleted WithdrawalStatus = "COMPLETED"
)

type OrderSide string

const (
	// OrderSideBuy represents buy orders.
	OrderSideBuy OrderSide = "BUY"

	// OrderSideSell represents sell orders.
	OrderSideSell OrderSide = "SELL"
)

type ChildOrderType string

const (
	// ChildOrderTypeLimit represents limit orders.
	ChildOrderTypeLimit ChildOrderType = "LIMIT"

	// ChildOrderTypeMarket represents market orders.
	ChildOrderTypeMarket ChildOrderType = "MARKET"
)

type ChildOrderState string

const (
	// ChildOrderStateActive represents open orders.
	ChildOrderStateActive ChildOrderState = "ACTIVE"

	// ChildOrderStateCompleted represents fully completed orders.
	ChildOrderStateCompleted ChildOrderState = "COMPLETED"

	// ChildOrderStateCancelled represents orders that have been cancelled by the customer.
	ChildOrderStateCancelled ChildOrderState = "CANCELED"

	// ChildOrderStateExpired represents orders that have been cancelled due to expiry.
	ChildOrderStateExpired ChildOrderState = "EXPIRED"

	// ChildOrderStateRejected represents failed orders.
	ChildOrderStateRejected ChildOrderState = "REJECTED"
)

type ParentOrderState string

const (
	// ParentOrderStateActive represents open orders.
	ParentOrderStateActive ParentOrderState = "ACTIVE"

	// ParentOrderStateCompleted represents fully completed orders.
	ParentOrderStateCompleted ParentOrderState = "COMPLETED"

	// ParentOrderStateCancelled represents orders that have been cancelled by the customer.
	ParentOrderStateCancelled ParentOrderState = "CANCELED"

	// ParentOrderStateExpired represents orders that have been cancelled due to expiry.
	ParentOrderStateExpired ParentOrderState = "EXPIRED"

	// ParentOrderStateRejected represents failed orders.
	ParentOrderStateRejected ParentOrderState = "REJECTED"
)

type TimeInForce string

const (
	// TimeInForceGTC represents A Good 'Til Canceled order that is one where the order remains in effect until it is either filled or canceled.
	TimeInForceGTC TimeInForce = "GTC"

	// TimeInForceIOC represents Immediate or Cancel order.
	TimeInForceIOC TimeInForce = "IOC"

	// TimeInForceFOK represents Fill or Kill that refer to when the order is canceled if the volume is not immediately contracted (filled) in its entirety.
	TimeInForceFOK TimeInForce = "FOK"
)
