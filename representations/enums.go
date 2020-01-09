package representations

type DecisionStrategy string

const (
	DecisionStrategy_Affirmative DecisionStrategy = "AFFIRMATIVE"
	DecisionStrategy_Unanimous   DecisionStrategy = "UNANIMOUS"
	DecisionStrategy_Consensus   DecisionStrategy = "CONSENSUS"
)

type PolicyEnforcementMode string

const (
	PolicyEnforcementMode_Enforcing  PolicyEnforcementMode = "ENFORCING"
	PolicyEnforcementMode_Permissive PolicyEnforcementMode = "PERMISSIVE"
	PolicyEnforcementMode_Disabled   PolicyEnforcementMode = "DISABLED"
)

type Logic string

const (
	Logic_Positive Logic = "POSITIVE"
	Logic_Negative Logic = "NEGATIVE"
)
