package types

type MarketType string

const (
	MarketTypeSpot    MarketType = "Spot"
	MarketTypeFx      MarketType = "FX"
	MarketTypeFutures MarketType = "Futures"
)

type Health string

const (
	// The exchange is operating.
	HealthNormal Health = "NORMAL"

	// The exchange is experiencing high traffic.
	HealthBusy Health = "BUSY"

	// The exchange is experiencing very heavy traffic.
	HealthVeryBusy Health = "VERY BUSY"

	// The exchange is experiencing extremely heavy traffic. There is a possibility that orders will fail or be processed after a delay.
	HealthSuperBusy Health = "SUPER BUSY"

	// Orders can not be received.
	HealthNoOrder Health = "NO ORDER"

	// The exchange has been stopped. Orders will not be accepted.
	HealthStop Health = "STOP"
)

type State string

const (
	// Operating
	StateRunning State = "RUNNING"

	// Suspending
	StateClosed State = "CLOSED"

	// Restarting
	StateStarting State = "STARTING"

	// Performing Itayose
	StatePreOpen State = "PREOPEN"

	// Circuit breaker triggered
	StateCIRCUITBREAK State = "CIRCUIT BREAK"

	// Calculating SQ (special quotation) for Lightning Futures after trades complete
	StateAwaitingSQ State = "AWAITING SQ"

	// Lightning Futures maturity reached
	StateMatured State = "MATURED"
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
