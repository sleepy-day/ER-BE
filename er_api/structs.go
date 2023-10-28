package er_api

type ERApiManager struct {
	apiKey string
}

// Response structs

type FreeCharacters struct {
	Code         int    `json:"code"`
	Message      string `json:"message"`
	CharacterIds []int  `json:"data"`
}

type LanguageUrl struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Path string `json:"l10Path"`
	} `json:"data"`
}

type ActionCostResponse struct {
	Code        int          `json:"code"`
	Message     string       `json:"message"`
	ActionCosts []ActionCost `json:"data"`
}

type AreaResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Areas   []Area `json:"data"`
}

type BattleZoneRewardResponse struct {
	Code              int                `json:"code"`
	Message           string             `json:"message"`
	BattleZoneRewards []BattleZoneReward `json:"data"`
}

type BulletCapacityResponse struct {
	Code           int              `json:"code"`
	Message        string           `json:"message"`
	BulletCapacity []BulletCapacity `json:"data"`
}

type CharacterResponse struct {
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Characters []Character `json:"data"`
}

type CharacterAttributesResponse struct {
	Code                int                   `json:"code"`
	Message             string                `json:"message"`
	CharacterAttributes []CharacterAttributes `json:"data"`
}

type CharacterLevelUpStatResponse struct {
	Code                  int                    `json:"code"`
	Message               string                 `json:"message"`
	CharacterLevelUpStats []CharacterLevelUpStat `json:"data"`
}

type CharacterMasteryResponse struct {
	Code             int                `json:"code"`
	Message          string             `json:"message"`
	CharacterMastery []CharacterMastery `json:"data"`
}

type CharacterModeModifierResponse struct {
	Code                  int                     `json:"code"`
	Message               string                  `json:"message"`
	CharacterModeModifier []CharacterModeModifier `json:"data"`
}

type CharacterSkinResponse struct {
	Code           int             `json:"code"`
	Message        string          `json:"message"`
	CharacterSkins []CharacterSkin `json:"data"`
}

type CollectibleResponse struct {
	Code         int           `json:"code"`
	Message      string        `json:"message"`
	Collectibles []Collectible `json:"data"`
}

type DropGroupResponse struct {
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	DropGroups []DropGroup `json:"data"`
}

type GainExpResponse struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	GainExp []GainExp `json:"data"`
}

type GainScoreResponse struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	GainScore []GainScore `json:"data"`
}

type GameTipResponse struct {
	Code     int       `json:"code"`
	Message  string    `json:"message"`
	GameTips []GameTip `json:"data"`
}

type HowToFindItemResponse struct {
	Code          int             `json:"code"`
	Message       string          `json:"message"`
	HowToFindItem []HowToFindItem `json:"data"`
}

type InfusionProductResponse struct {
	Code             int               `json:"code"`
	Message          string            `json:"message"`
	InfusionProducts []InfusionProduct `json:"data"`
}

type ItemArmorResponse struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	ItemArmor []ItemArmor `json:"data"`
}

type ItemConsumableResponse struct {
	Code           int              `json:"code"`
	Message        string           `json:"message"`
	ItemConsumable []ItemConsumable `json:"data"`
}

type ItemMiscResponse struct {
	Code     int        `json:"code"`
	Message  string     `json:"message"`
	ItemMisc []ItemMisc `json:"data"`
}

type ItemSearchOptionResponse struct {
	Code             int                `json:"code"`
	Message          string             `json:"message"`
	ItemSearchOption []ItemSearchOption `json:"data"`
}

type ItemSpawnResponse struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	ItemSpawn []ItemSpawn `json:"data"`
}

type ItemSpecialResponse struct {
	Code        int           `json:"code"`
	Message     string        `json:"message"`
	ItemSpecial []ItemSpecial `json:"data"`
}

type ItemWeaponResponse struct {
	Code       int          `json:"code"`
	Message    string       `json:"message"`
	ItemWeapon []ItemWeapon `json:"data"`
}

type LevelResponse struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Level   []Level `json:"data"`
}

type LoadingTipResponse struct {
	Code       int          `json:"code"`
	Message    string       `json:"message"`
	LoadingTip []LoadingTip `json:"data"`
}

type MasteryExpResponse struct {
	Code       int          `json:"code"`
	Message    string       `json:"message"`
	MasteryExp []MasteryExp `json:"data"`
}

type MasteryLevelResponse struct {
	Code         int            `json:"code"`
	Message      string         `json:"message"`
	MasteryLevel []MasteryLevel `json:"data"`
}

type MasteryStatResponse struct {
	Code        int           `json:"code"`
	Message     string        `json:"message"`
	MasteryStat []MasteryStat `json:"data"`
}

type MonsterResponse struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Monster []Monster `json:"data"`
}

type MonsterDropGroupResponse struct {
	Code             int                `json:"code"`
	Message          string             `json:"message"`
	MonsterDropGroup []MonsterDropGroup `json:"data"`
}

type MonsterLevelUpStatResponse struct {
	Code               int                  `json:"code"`
	Message            string               `json:"message"`
	MonsterLevelUpStat []MonsterLevelUpStat `json:"data"`
}

type MonsterSpawnLevelResponse struct {
	Code              int                 `json:"code"`
	Message           string              `json:"message"`
	MonsterSpawnLevel []MonsterSpawnLevel `json:"data"`
}

type NaviCollectAndHuntResponse struct {
	Code               int                  `json:"code"`
	Message            string               `json:"message"`
	NaviCollectAndHunt []NaviCollectAndHunt `json:"data"`
}

type NearByAreaResponse struct {
	Code       int          `json:"code"`
	Message    string       `json:"message"`
	NearByArea []NearByArea `json:"data"`
}

type RandomEquipmentResponse struct {
	Code            int               `json:"code"`
	Message         string            `json:"message"`
	RandomEquipment []RandomEquipment `json:"data"`
}

type RecommendedListResponse struct {
	Code            int               `json:"code"`
	Message         string            `json:"message"`
	RecommendedList []RecommendedList `json:"data"`
}

type SeasonResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Season  []Season `json:"data"`
}

type SummonObjectStatResponse struct {
	Code             int                `json:"code"`
	Message          string             `json:"message"`
	SummonObjectStat []SummonObjectStat `json:"data"`
}

type TacticalSkillSetResponse struct {
	Code             int                `json:"code"`
	Message          string             `json:"message"`
	TacticalSkillSet []TacticalSkillSet `json:"data"`
}

type TacticalSkillSetGroupResponse struct {
	Code                  int                     `json:"code"`
	Message               string                  `json:"message"`
	TacticalSkillSetGroup []TacticalSkillSetGroup `json:"data"`
}

type TraitResponse struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Trait   []Trait `json:"data"`
}

type TransferConsoleResponse struct {
	Code            int               `json:"code"`
	Message         string            `json:"message"`
	TransferConsole []TransferConsole `json:"data"`
}

type VFCreditResponse struct {
	Code     int        `json:"code"`
	Message  string     `json:"message"`
	VFCredit []VFCredit `json:"data"`
}

type WeaponTypeInfoResponse struct {
	Code           int              `json:"code"`
	Message        string           `json:"message"`
	WeaponTypeInfo []WeaponTypeInfo `json:"data"`
}

type GameInfoResponse struct {
	Code     int        `json:"code"`
	Message  string     `json:"message"`
	GameInfo []GameInfo `json:"userGames"`
	Next     int        `json:"next,omitempty"`
}

type TopRankedPlayersResponse struct {
	Code     int        `json:"code"`
	Message  string     `json:"message"`
	GameInfo []GameInfo `json:"topRanks"`
}

type UserRankInfoResponse struct {
	Code         int          `json:"code"`
	Message      string       `json:"message"`
	UserRankInfo UserRankInfo `json:"userRank"`
}

type UserInfoResponse struct {
	Code     int      `json:"code"`
	Message  string   `json:"message"`
	UserInfo UserInfo `json:"user"`
}

type UserStatsResponse struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	UserStats []UserStats `json:"userStats"`
}

type ItemBuildsResponse struct {
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	ItemBuilds []ItemBuild `json:"result"`
}

type ItemBuildResponse struct {
	Code      int       `json:"code"`
	Message   string    `json:"message"`
	ItemBuild ItemBuild `json:"result"`
}

// Game data structs

type CharacterExpResponse struct {
	Code         int            `json:"code"`
	Message      string         `json:"message"`
	CharacterExp []CharacterExp `json:"data"`
}

type Area struct {
	Code                     int    `json:"code"`
	Name                     string `json:"name"`
	ModeType                 int    `json:"modeType"`
	MaskCode                 int    `json:"maskCode"`
	StartingArea             bool   `json:"startingArea"`
	AreaType                 string `json:"areaType"`
	IsProvideCollectibleItem bool   `json:"isProvideCollectibleItem"`
	RouteCalcBitCode         int    `json:"routeCalcBitCode"`
	IsHyperLoopInstalled     bool   `json:"isHyperLoopInstalled"`
}

type ActionCost struct {
	Code                  int     `json:"code"`
	Type                  string  `json:"type"`
	Sp                    int     `json:"sp"`
	InitialTime           float32 `json:"time1"`
	AdditionalTime        float32 `json:"time2"`
	ActionWaitTime        float32 `json:"actionWaitTime"`
	CastingAnimTrigger    string  `json:"castingAnimTrigger"`
	EffectCancelCondition string  `json:"effectCancelCondition"`
	CastingBarImgName     string  `json:"castingBarImgName"`
}

type BattleZoneReward struct {
	Code                           int    `json:"code"`
	ModeType                       int    `json:"modeType"`
	AreaAttributesCreateEventCount int    `json:"areaAttributesCreateEventCount"`
	ItemCode                       int    `json:"itemCode"`
	Type                           string `json:"type"`
	Value                          int    `json:"value"`
	Selectable                     bool   `json:"selectable"`
}

type BulletCapacity struct {
	ItemCode  int    `json:"itemCode"`
	Capacity  int    `json:"capacity"`
	LoadType  string `json:"loadType"`
	Time      int    `json:"time"`
	InitCount int    `json:"initCount"`
	Count     int    `json:"count"`
}

type Character struct {
	Code                           int     `json:"code"`
	Name                           string  `json:"name"`
	MaxHp                          int     `json:"maxHp"`
	MaxSp                          int     `json:"maxSp"`
	StrLearnStartSkill             string  `json:"strLearnStartSkill"`
	StrUsePointLearnStartSkill     string  `json:"strUsePointLearnStartSkill"`
	InitExtraPoint                 int     `json:"initExtraPoint"`
	MaxExtraPoint                  int     `json:"maxExtraPoint"`
	AttackPower                    int     `json:"attackPower"`
	Defense                        int     `json:"defense"`
	SkillAmp                       int     `json:"skillAmp"`
	AdaptiveForce                  int     `json:"adaptiveForce"`
	CriticalStrikeChance           int     `json:"criticalStrikeChance"`
	HpRegen                        float64 `json:"hpRegen"`
	SpRegen                        float64 `json:"spRegen"`
	AttackSpeed                    float64 `json:"attackSpeed"`
	AttackSpeedRatio               int     `json:"attackSpeedRatio"`
	IncreaseBasicAttackDamageRatio int     `json:"increaseBasicAttackDamageRatio"`
	SkillAmpRatio                  int     `json:"skillAmpRatio"`
	PreventBasicAttackDamagedRatio int     `json:"preventBasicAttackDamagedRatio"`
	PreventSkillDamagedRatio       int     `json:"preventSkillDamagedRatio"`
	AttackSpeedLimit               float64 `json:"attackSpeedLimit"`
	AttackSpeedMin                 int     `json:"attackSpeedMin"`
	MoveSpeed                      float64 `json:"moveSpeed"`
	SightRange                     float64 `json:"sightRange"`
	Radius                         float64 `json:"radius"`
	PathingRadius                  float64 `json:"pathingRadius"`
	UIHeight                       float64 `json:"uiHeight"`
	InitStateDisplayIndex          int     `json:"initStateDisplayIndex"`
	LocalScaleInCutscene           int     `json:"localScaleInCutscene"`
	LocalScaleInVictoryScene       string  `json:"localScaleInVictoryScene"`
	Resource                       string  `json:"resource"`
	LobbySubObject                 string  `json:"lobbySubObject"`
}

type CharacterAttributes struct {
	Character         string `json:"character"`
	CharacterCode     int    `json:"characterCode"`
	Mastery           string `json:"mastery"`
	ControlDifficulty int    `json:"controlDifficulty"`
	Attack            int    `json:"attack"`
	Defense           int    `json:"defense"`
	Disruptor         int    `json:"disruptor"`
	Move              int    `json:"move"`
	Assistance        int    `json:"assistance"`
}

type CharacterExp struct {
	Level      int `json:"level"`
	LevelUpExp int `json:"levelUpExp"`
}

type CharacterLevelUpStat struct {
	Code                           int     `json:"code"`
	Name                           string  `json:"name"`
	MaxHp                          int     `json:"maxHp"`
	MaxSp                          int     `json:"maxSp"`
	AttackPower                    float64 `json:"attackPower"`
	Defense                        float64 `json:"defense"`
	SkillAmp                       int     `json:"skillAmp"`
	AdaptiveForce                  int     `json:"adaptiveForce"`
	CriticalChance                 int     `json:"criticalChance"`
	HpRegen                        float64 `json:"hpRegen"`
	SpRegen                        float64 `json:"spRegen"`
	AttackSpeed                    int     `json:"attackSpeed"`
	MoveSpeed                      int     `json:"moveSpeed"`
	AttackSpeedRatio               int     `json:"attackSpeedRatio"`
	IncreaseBasicAttackDamageRatio int     `json:"increaseBasicAttackDamageRatio"`
	SkillAmpRatio                  int     `json:"skillAmpRatio"`
	PreventBasicAttackDamagedRatio int     `json:"preventBasicAttackDamagedRatio"`
	PreventSkillDamagedRatio       int     `json:"preventSkillDamagedRatio"`
}

type CharacterMastery struct {
	Code      int    `json:"code"`
	Weapon1   string `json:"weapon1"`
	Weapon2   string `json:"weapon2"`
	Weapon3   string `json:"weapon3"`
	Weapon4   string `json:"weapon4"`
	Combat1   string `json:"combat1"`
	Combat2   string `json:"combat2"`
	Survival1 string `json:"survival1"`
	Survival2 string `json:"survival2"`
	Survival3 string `json:"survival3"`
}

type CharacterModeModifier struct {
	CharacterCode                           int    `json:"characterCode"`
	WeaponType                              string `json:"weaponType"`
	SoloIncreaseModeDamageRatio             int    `json:"soloIncreaseModeDamageRatio"`
	SoloPreventModeDamageRatio              int    `json:"soloPreventModeDamageRatio"`
	SoloIncreaseModeHealRatio               int    `json:"soloIncreaseModeHealRatio"`
	SoloIncreaseModeShieldRatio             int    `json:"soloIncreaseModeShieldRatio"`
	DuoIncreaseModeDamageRatio              int    `json:"duoIncreaseModeDamageRatio"`
	DuoPreventModeDamageRatio               int    `json:"duoPreventModeDamageRatio"`
	DuoIncreaseModeHealRatio                int    `json:"duoIncreaseModeHealRatio"`
	DuoIncreaseModeHealerGiveHealRatio      int    `json:"duoIncreaseModeHealerGiveHealRatio"`
	DuoIncreaseModeShieldRatio              int    `json:"duoIncreaseModeShieldRatio"`
	DuoIncreaseModeHealerGiveShieldRatio    int    `json:"duoIncreaseModeHealerGiveShieldRatio"`
	SquadIncreaseModeDamageRatio            int    `json:"squadIncreaseModeDamageRatio"`
	SquadPreventModeDamageRatio             int    `json:"squadPreventModeDamageRatio"`
	SquadIncreaseModeHealRatio              int    `json:"squadIncreaseModeHealRatio"`
	SquadIncreaseModeHealerGiveHealRatio    int    `json:"squadIncreaseModeHealerGiveHealRatio"`
	SquadIncreaseModeShieldRatio            int    `json:"squadIncreaseModeShieldRatio"`
	SquadIncreaseModeHealerGiveShieldRatio  int    `json:"squadIncreaseModeHealerGiveShieldRatio"`
	CobaltIncreaseModeDamageRatio           int    `json:"cobaltIncreaseModeDamageRatio"`
	CobaltPreventModeDamageRatio            int    `json:"cobaltPreventModeDamageRatio"`
	CobaltIncreaseModeHealRatio             int    `json:"cobaltIncreaseModeHealRatio"`
	CobaltIncreaseModeHealerGiveHealRatio   int    `json:"cobaltIncreaseModeHealerGiveHealRatio"`
	CobaltIncreaseModeShieldRatio           int    `json:"cobaltIncreaseModeShieldRatio"`
	CobaltIncreaseModeHealerGiveShieldRatio int    `json:"cobaltIncreaseModeHealerGiveShieldRatio"`
	CobaltIncreaseModeUltCooldownRatio      int    `json:"cobaltIncreaseModeUltCooldownRatio"`
	CobaltIncreaseModeMaxSpRatio            int    `json:"cobaltIncreaseModeMaxSpRatio"`
	CobaltIncreaseModeSpRegenRatio          int    `json:"cobaltIncreaseModeSpRegenRatio"`
	SoloIncreaseModeDamageToMonsterRatio    int    `json:"soloIncreaseModeDamageToMonsterRatio"`
	DuoIncreaseModeDamageToMonsterRatio     int    `json:"duoIncreaseModeDamageToMonsterRatio"`
	SquadIncreaseModeDamageToMonsterRatio   int    `json:"squadIncreaseModeDamageToMonsterRatio"`
	CobaltIncreaseModeDamageToMonsterRatio  int    `json:"cobaltIncreaseModeDamageToMonsterRatio"`
}

type CharacterSkin struct {
	Name                     string `json:"name"`
	Code                     int    `json:"code"`
	CharacterCode            int    `json:"characterCode"`
	Index                    int    `json:"index"`
	Grade                    int    `json:"grade"`
	PurchaseType             string `json:"purchaseType"`
	EffectsPath              string `json:"effectsPath"`
	ProjectilesPath          string `json:"projectilesPath"`
	ObjectPath               string `json:"objectPath"`
	FxSoundPath              string `json:"fxSoundPath"`
	VoicePath                string `json:"voicePath"`
	WeaponMountPath          string `json:"weaponMountPath"`
	WeaponMountCommonPath    string `json:"weaponMountCommonPath"`
	IndicatorPath            string `json:"indicatorPath"`
	ProjectilesDeflectorPath string `json:"projectilesDeflectorPath"`
}

type Collectible struct {
	Code              int    `json:"code"`
	Cooldown          int    `json:"cooldown"`
	ItemCode1         string `json:"itemCode1"`
	ItemCode2         string `json:"itemCode2"`
	Probability1      int    `json:"probability1"`
	Probability2      int    `json:"probability2"`
	DropCount         int    `json:"dropCount"`
	CastingActionType string `json:"castingActionType"`
}

type DropGroup struct {
	GroupCode   int    `json:"groupCode"`
	ItemCode    string `json:"itemCode"`
	Min         int    `json:"min"`
	Max         int    `json:"max"`
	Probability int    `json:"probability"`
	DropType    string `json:"dropType"`
}

type GainExp struct {
	StartTime int `json:"startTime"`
	EndTime   int `json:"endTime"`
	GainExp   int `json:"gainExp"`
}

type GainScore struct {
	Code           int    `json:"code"`
	Phase          int    `json:"phase"`
	ConditionType  string `json:"conditionType"`
	ConditionValue int    `json:"conditionValue"`
	PointsEnemy    int    `json:"pointsEnemy"`
	PointsAlly     int    `json:"pointsAlly"`
}

type GameTip struct {
	Key            int    `json:"key"`
	Code           int    `json:"code"`
	GameTipType    string `json:"gameTipType"`
	Sequence       int    `json:"sequence"`
	TitleTextKey   string `json:"titleTextKey"`
	ContentTextKey string `json:"contentTextKey"`
	ImageName      string `json:"imageName"`
	Link           string `json:"link"`
}

type HowToFindItem struct {
	Code            int `json:"code"`
	ItemCode        int `json:"itemCode"`
	HuntChicken     int `json:"huntChicken"`
	HuntBat         int `json:"huntBat"`
	HuntBoar        int `json:"huntBoar"`
	HuntWildDog     int `json:"huntWildDog"`
	HuntWolf        int `json:"huntWolf"`
	HuntBear        int `json:"huntBear"`
	HuntWickline    int `json:"huntWickline"`
	HuntAlpha       int `json:"huntAlpha"`
	HuntOmega       int `json:"huntOmega"`
	CollectibleCode int `json:"collectibleCode"`
	AirSupply       int `json:"airSupply"`
}

type InfusionProduct struct {
	Code             int    `json:"code"`
	ProductType      string `json:"productType"`
	ProductGroup     int    `json:"productGroup"`
	ProductCode      int    `json:"productCode"`
	StoreType        string `json:"storeType"`
	StockType        string `json:"stockType"`
	Stock            int    `json:"stock"`
	IsRestore        bool   `json:"isRestore"`
	Price            int    `json:"price"`
	SpecialWeight    int    `json:"specialWeight"`
	Weight           int    `json:"weight"`
	Requirement      int    `json:"requirement"`
	Icon             string `json:"icon"`
	SimpleIcon       string `json:"simpleIcon"`
	AlertInSpectator bool   `json:"alertInSpectator"`
	CharacterNum     int    `json:"characterNum"`
}

type ItemArmor struct {
	Code                                               int    `json:"code"`
	Name                                               string `json:"name"`
	ModeType                                           int    `json:"modeType"`
	ItemType                                           string `json:"itemType"`
	ArmorType                                          string `json:"armorType"`
	ItemGrade                                          string `json:"itemGrade"`
	IsCompletedItem                                    bool   `json:"isCompletedItem"`
	AlertInSpectator                                   bool   `json:"alertInSpectator"`
	MarkingType                                        string `json:"markingType"`
	CraftAnimTrigger                                   string `json:"craftAnimTrigger"`
	Stackable                                          int    `json:"stackable"`
	InitialCount                                       int    `json:"initialCount"`
	ItemUsableType                                     string `json:"itemUsableType"`
	ItemUsableValueList                                int    `json:"itemUsableValueList"`
	ExclusiveProducer                                  int    `json:"exclusiveProducer"`
	IsRemovedFromPlayerCorpseInventoryWhenPlayerKilled bool   `json:"isRemovedFromPlayerCorpseInventoryWhenPlayerKilled"`
	MakeMaterial1                                      int    `json:"makeMaterial1"`
	MakeMaterial2                                      int    `json:"makeMaterial2"`
	NotDisarm                                          bool   `json:"notDisarm"`
	ManufacturableType                                 int    `json:"manufacturableType"`
	AttackPower                                        int    `json:"attackPower"`
	AttackPowerByLv                                    int    `json:"attackPowerByLv"`
	Defense                                            int    `json:"defense"`
	DefenseByLv                                        int    `json:"defenseByLv"`
	SkillAmp                                           int    `json:"skillAmp"`
	SkillAmpByLevel                                    int    `json:"skillAmpByLevel"`
	SkillAmpRatio                                      int    `json:"skillAmpRatio"`
	SkillAmpRatioByLevel                               int    `json:"skillAmpRatioByLevel"`
	AdaptiveForce                                      int    `json:"adaptiveForce"`
	AdaptiveForceByLevel                               int    `json:"adaptiveForceByLevel"`
	MaxHp                                              int    `json:"maxHp"`
	MaxHpByLv                                          int    `json:"maxHpByLv"`
	MaxSp                                              int    `json:"maxSp"`
	HpRegenRatio                                       int    `json:"hpRegenRatio"`
	HpRegen                                            int    `json:"hpRegen"`
	SpRegenRatio                                       int    `json:"spRegenRatio"`
	SpRegen                                            int    `json:"spRegen"`
	AttackSpeedRatio                                   int    `json:"attackSpeedRatio"`
	AttackSpeedRatioByLv                               int    `json:"attackSpeedRatioByLv"`
	CriticalStrikeChance                               int    `json:"criticalStrikeChance"`
	CriticalStrikeDamage                               int    `json:"criticalStrikeDamage"`
	PreventCriticalStrikeDamaged                       int    `json:"preventCriticalStrikeDamaged"`
	CooldownReduction                                  int    `json:"cooldownReduction"`
	CooldownLimit                                      int    `json:"cooldownLimit"`
	LifeSteal                                          int    `json:"lifeSteal"`
	NormalLifeSteal                                    int    `json:"normalLifeSteal"`
	SkillLifeSteal                                     int    `json:"skillLifeSteal"`
	MoveSpeed                                          int    `json:"moveSpeed"`
	MoveSpeedOutOfCombat                               int    `json:"moveSpeedOutOfCombat"`
	SightRange                                         int    `json:"sightRange"`
	AttackRange                                        int    `json:"attackRange"`
	IncreaseBasicAttackDamage                          int    `json:"increaseBasicAttackDamage"`
	IncreaseBasicAttackDamageByLv                      int    `json:"increaseBasicAttackDamageByLv"`
	PreventBasicAttackDamaged                          int    `json:"preventBasicAttackDamaged"`
	PreventBasicAttackDamagedByLv                      int    `json:"preventBasicAttackDamagedByLv"`
	PreventBasicAttackDamagedRatio                     int    `json:"preventBasicAttackDamagedRatio"`
	PreventBasicAttackDamagedRatioByLv                 int    `json:"preventBasicAttackDamagedRatioByLv"`
	IncreaseBasicAttackDamageRatio                     int    `json:"increaseBasicAttackDamageRatio"`
	IncreaseBasicAttackDamageRatioByLv                 int    `json:"increaseBasicAttackDamageRatioByLv"`
	PreventSkillDamaged                                int    `json:"preventSkillDamaged"`
	PreventSkillDamagedByLv                            int    `json:"preventSkillDamagedByLv"`
	PreventSkillDamagedRatio                           int    `json:"preventSkillDamagedRatio"`
	PreventSkillDamagedRatioByLv                       int    `json:"preventSkillDamagedRatioByLv"`
	PenetrationDefense                                 int    `json:"penetrationDefense"`
	PenetrationDefenseRatio                            int    `json:"penetrationDefenseRatio"`
	TrapDamageReduce                                   int    `json:"trapDamageReduce"`
	TrapDamageReduceRatio                              int    `json:"trapDamageReduceRatio"`
	HpHealedIncreaseRatio                              int    `json:"hpHealedIncreaseRatio"`
	HealerGiveHpHealRatio                              int    `json:"healerGiveHpHealRatio"`
	UniqueAttackRange                                  int    `json:"uniqueAttackRange"`
	UniqueHpHealedIncreaseRatio                        int    `json:"uniqueHpHealedIncreaseRatio"`
	UniqueCooldownLimit                                int    `json:"uniqueCooldownLimit"`
	UniqueTenacity                                     int    `json:"uniqueTenacity"`
	UniqueMoveSpeed                                    int    `json:"uniqueMoveSpeed"`
	UniquePenetrationDefense                           int    `json:"uniquePenetrationDefense"`
	UniquePenetrationDefenseRatio                      int    `json:"uniquePenetrationDefenseRatio"`
	UniqueLifeSteal                                    int    `json:"uniqueLifeSteal"`
	UniqueSkillAmpRatio                                int    `json:"uniqueSkillAmpRatio"`
	RestoreItemWhenResurrected                         bool   `json:"restoreItemWhenResurrected"`
	CreditValueWhenConvertedToBounty                   int    `json:"creditValueWhenConvertedToBounty"`
}

type ItemConsumable struct {
	Code                             int    `json:"code"`
	Name                             string `json:"name"`
	ModeType                         int    `json:"modeType"`
	ItemType                         string `json:"itemType"`
	ConsumableType                   string `json:"consumableType"`
	ConsumableTag                    string `json:"consumableTag"`
	ItemGrade                        string `json:"itemGrade"`
	IsCompletedItem                  bool   `json:"isCompletedItem"`
	AlertInSpectator                 bool   `json:"alertInSpectator"`
	MarkingType                      string `json:"markingType"`
	CraftAnimTrigger                 string `json:"craftAnimTrigger"`
	Stackable                        int    `json:"stackable"`
	InitialCount                     int    `json:"initialCount"`
	ItemUsableType                   string `json:"itemUsableType"`
	ItemUsableValueList              int    `json:"itemUsableValueList"`
	ExclusiveProducer                int    `json:"exclusiveProducer"`
	CanNotBeTakeItemFromCorpse       bool   `json:"canNotBeTakeItemFromCorpse"`
	ManufacturableType               int    `json:"manufacturableType"`
	MakeMaterial1                    int    `json:"makeMaterial1"`
	MakeMaterial2                    int    `json:"makeMaterial2"`
	Heal                             int    `json:"heal"`
	HpRecover                        int    `json:"hpRecover"`
	SpRecover                        int    `json:"spRecover"`
	AttackPowerByBuff                int    `json:"attackPowerByBuff"`
	DefenseByBuff                    int    `json:"defenseByBuff"`
	SkillAmpByBuff                   int    `json:"skillAmpByBuff"`
	SkillAmpRatioByBuff              int    `json:"skillAmpRatioByBuff"`
	AddStateCode                     int    `json:"addStateCode"`
	IsVPadQuickSlotItem              bool   `json:"isVPadQuickSlotItem"`
	RestoreItemWhenResurrected       bool   `json:"restoreItemWhenResurrected"`
	CreditValueWhenConvertedToBounty int    `json:"creditValueWhenConvertedToBounty"`
}

type ItemMisc struct {
	Code                             int    `json:"code"`
	Name                             string `json:"name"`
	ModeType                         int    `json:"modeType"`
	ItemType                         string `json:"itemType"`
	MiscItemType                     string `json:"miscItemType"`
	ItemGrade                        string `json:"itemGrade"`
	IsCompletedItem                  bool   `json:"isCompletedItem"`
	AlertInSpectator                 bool   `json:"alertInSpectator"`
	MarkingType                      string `json:"markingType"`
	CraftAnimTrigger                 string `json:"craftAnimTrigger"`
	Stackable                        int    `json:"stackable"`
	InitialCount                     int    `json:"initialCount"`
	ItemUsableType                   string `json:"itemUsableType"`
	ItemUsableValueList              int    `json:"itemUsableValueList"`
	ExclusiveProducer                int    `json:"exclusiveProducer"`
	CanNotBeTakeItemFromCorpse       bool   `json:"canNotBeTakeItemFromCorpse"`
	ManufacturableType               int    `json:"manufacturableType"`
	MakeMaterial1                    int    `json:"makeMaterial1"`
	MakeMaterial2                    int    `json:"makeMaterial2"`
	RestoreItemWhenResurrected       bool   `json:"restoreItemWhenResurrected"`
	CreditValueWhenConvertedToBounty int    `json:"creditValueWhenConvertedToBounty"`
}

type ItemSearchOption struct {
	Code                       int    `json:"code"`
	Name                       string `json:"name"`
	AttackPowerSearch          bool   `json:"attackPowerSearch"`
	AttackSpeedSearch          bool   `json:"attackSpeedSearch"`
	CriticalStrikeSearch       bool   `json:"criticalStrikeSearch"`
	LifeStealSearch            bool   `json:"lifeStealSearch"`
	NormalAttackIncreaseSearch bool   `json:"normalAttackIncreaseSearch"`
	SkillAttackIncreaseSearch  bool   `json:"skillAttackIncreaseSearch"`
	CooldownReductionSearch    bool   `json:"cooldownReductionSearch"`
	MaxSpSearch                bool   `json:"maxSpSearch"`
	SpRegenSearch              bool   `json:"spRegenSearch"`
	DefenseSearch              bool   `json:"defenseSearch"`
	MaxHpSearch                bool   `json:"maxHpSearch"`
	HpRegenSearch              bool   `json:"hpRegenSearch"`
	DamageReductionSearch      bool   `json:"damageReductionSearch"`
	MoveSpeedSearch            bool   `json:"moveSpeedSearch"`
	SightRangeSearch           bool   `json:"sightRangeSearch"`
	Tag1                       string `json:"tag1"`
	Tag2                       string `json:"tag2"`
	Tag3                       string `json:"tag3"`
}

type ItemSpawn struct {
	Code           int    `json:"code"`
	AreaCode       int    `json:"areaCode"`
	AreaSpawnGroup int    `json:"areaSpawnGroup"`
	ItemCode       int    `json:"itemCode"`
	DropPoint      string `json:"dropPoint"`
	DropCount      int    `json:"dropCount"`
}

type ItemSpecial struct {
	Code                             int    `json:"code"`
	Name                             string `json:"name"`
	ModeType                         int    `json:"modeType"`
	ItemType                         string `json:"itemType"`
	SpecialItemType                  string `json:"specialItemType"`
	ItemGrade                        string `json:"itemGrade"`
	IsCompletedItem                  bool   `json:"isCompletedItem"`
	AlertInSpectator                 bool   `json:"alertInSpectator"`
	MarkingType                      string `json:"markingType"`
	CraftAnimTrigger                 string `json:"craftAnimTrigger"`
	Stackable                        int    `json:"stackable"`
	InitialCount                     int    `json:"initialCount"`
	CooldownGroupCode                int    `json:"cooldownGroupCode"`
	Cooldown                         int    `json:"cooldown"`
	ItemUsableType                   string `json:"itemUsableType"`
	ItemUsableValueList              int    `json:"itemUsableValueList"`
	ExclusiveProducer                int    `json:"exclusiveProducer"`
	CanNotBeTakeItemFromCorpse       bool   `json:"canNotBeTakeItemFromCorpse"`
	ManufacturableType               int    `json:"manufacturableType"`
	MakeMaterial1                    int    `json:"makeMaterial1"`
	MakeMaterial2                    int    `json:"makeMaterial2"`
	ConsumeCount                     int    `json:"consumeCount"`
	SummonCode                       int    `json:"summonCode"`
	GhostItemStateGroup              int    `json:"ghostItemStateGroup"`
	IsVPadQuickSlotItem              bool   `json:"isVPadQuickSlotItem"`
	RestoreItemWhenResurrected       bool   `json:"restoreItemWhenResurrected"`
	CreditValueWhenConvertedToBounty int    `json:"creditValueWhenConvertedToBounty"`
}

type ItemWeapon struct {
	Code                                               int     `json:"code"`
	Name                                               string  `json:"name"`
	ModeType                                           int     `json:"modeType"`
	ItemType                                           string  `json:"itemType"`
	WeaponType                                         string  `json:"weaponType"`
	ItemGrade                                          string  `json:"itemGrade"`
	IsCompletedItem                                    bool    `json:"isCompletedItem"`
	AlertInSpectator                                   bool    `json:"alertInSpectator"`
	MarkingType                                        string  `json:"markingType"`
	CraftAnimTrigger                                   string  `json:"craftAnimTrigger"`
	Stackable                                          int     `json:"stackable"`
	InitialCount                                       int     `json:"initialCount"`
	ItemUsableType                                     string  `json:"itemUsableType"`
	ItemUsableValueList                                int     `json:"itemUsableValueList"`
	ExclusiveProducer                                  int     `json:"exclusiveProducer"`
	IsRemovedFromPlayerCorpseInventoryWhenPlayerKilled bool    `json:"isRemovedFromPlayerCorpseInventoryWhenPlayerKilled"`
	MakeMaterial1                                      int     `json:"makeMaterial1"`
	MakeMaterial2                                      int     `json:"makeMaterial2"`
	NotDisarm                                          bool    `json:"notDisarm"`
	Consumable                                         bool    `json:"consumable"`
	ManufacturableType                                 int     `json:"manufacturableType"`
	AttackPower                                        int     `json:"attackPower"`
	AttackPowerByLv                                    int     `json:"attackPowerByLv"`
	Defense                                            int     `json:"defense"`
	DefenseByLv                                        int     `json:"defenseByLv"`
	SkillAmp                                           int     `json:"skillAmp"`
	SkillAmpByLevel                                    int     `json:"skillAmpByLevel"`
	SkillAmpRatio                                      int     `json:"skillAmpRatio"`
	SkillAmpRatioByLevel                               int     `json:"skillAmpRatioByLevel"`
	AdaptiveForce                                      int     `json:"adaptiveForce"`
	AdaptiveForceByLevel                               int     `json:"adaptiveForceByLevel"`
	MaxHp                                              int     `json:"maxHp"`
	MaxHpByLv                                          int     `json:"maxHpByLv"`
	HpRegenRatio                                       int     `json:"hpRegenRatio"`
	HpRegen                                            int     `json:"hpRegen"`
	MaxSP                                              int     `json:"maxSP"`
	SpRegenRatio                                       int     `json:"spRegenRatio"`
	SpRegen                                            int     `json:"spRegen"`
	AttackSpeedRatio                                   int     `json:"attackSpeedRatio"`
	AttackSpeedRatioByLv                               int     `json:"attackSpeedRatioByLv"`
	CriticalStrikeChance                               int     `json:"criticalStrikeChance"`
	CriticalStrikeDamage                               int     `json:"criticalStrikeDamage"`
	CooldownReduction                                  int     `json:"cooldownReduction"`
	PreventCriticalStrikeDamaged                       int     `json:"preventCriticalStrikeDamaged"`
	CooldownLimit                                      int     `json:"cooldownLimit"`
	LifeSteal                                          int     `json:"lifeSteal"`
	NormalLifeSteal                                    int     `json:"normalLifeSteal"`
	SkillLifeSteal                                     int     `json:"skillLifeSteal"`
	MoveSpeed                                          float64 `json:"moveSpeed"`
	MoveSpeedOutOfCombat                               int     `json:"moveSpeedOutOfCombat"`
	SightRange                                         int     `json:"sightRange"`
	AttackRange                                        int     `json:"attackRange"`
	IncreaseBasicAttackDamage                          int     `json:"increaseBasicAttackDamage"`
	IncreaseBasicAttackDamageByLv                      int     `json:"increaseBasicAttackDamageByLv"`
	IncreaseBasicAttackDamageRatio                     int     `json:"increaseBasicAttackDamageRatio"`
	IncreaseBasicAttackDamageRatioByLv                 int     `json:"increaseBasicAttackDamageRatioByLv"`
	PreventBasicAttackDamaged                          int     `json:"preventBasicAttackDamaged"`
	PreventBasicAttackDamagedByLv                      int     `json:"preventBasicAttackDamagedByLv"`
	PreventBasicAttackDamagedRatio                     int     `json:"preventBasicAttackDamagedRatio"`
	PreventBasicAttackDamagedRatioByLv                 int     `json:"preventBasicAttackDamagedRatioByLv"`
	PreventSkillDamaged                                int     `json:"preventSkillDamaged"`
	PreventSkillDamagedByLv                            int     `json:"preventSkillDamagedByLv"`
	PreventSkillDamagedRatio                           int     `json:"preventSkillDamagedRatio"`
	PreventSkillDamagedRatioByLv                       int     `json:"preventSkillDamagedRatioByLv"`
	PenetrationDefense                                 int     `json:"penetrationDefense"`
	PenetrationDefenseRatio                            int     `json:"penetrationDefenseRatio"`
	TrapDamageReduce                                   int     `json:"trapDamageReduce"`
	TrapDamageReduceRatio                              int     `json:"trapDamageReduceRatio"`
	HpHealedIncreaseRatio                              int     `json:"hpHealedIncreaseRatio"`
	HealerGiveHpHealRatio                              int     `json:"healerGiveHpHealRatio"`
	UniqueAttackRange                                  int     `json:"uniqueAttackRange"`
	UniqueHpHealedIncreaseRatio                        int     `json:"uniqueHpHealedIncreaseRatio"`
	UniqueCooldownLimit                                int     `json:"uniqueCooldownLimit"`
	UniqueTenacity                                     int     `json:"uniqueTenacity"`
	UniqueMoveSpeed                                    int     `json:"uniqueMoveSpeed"`
	UniquePenetrationDefense                           int     `json:"uniquePenetrationDefense"`
	UniquePenetrationDefenseRatio                      int     `json:"uniquePenetrationDefenseRatio"`
	UniqueLifeSteal                                    int     `json:"uniqueLifeSteal"`
	UniqueSkillAmpRatio                                int     `json:"uniqueSkillAmpRatio"`
	RestoreItemWhenResurrected                         bool    `json:"restoreItemWhenResurrected"`
	CreditValueWhenConvertedToBounty                   int     `json:"creditValueWhenConvertedToBounty"`
}

type Level struct {
	Level         int `json:"level"`
	NeedExp       int `json:"needExp"`
	AccumulateExp int `json:"accumulateExp"`
	RewardAcoin   int `json:"rewardAcoin"`
	Reward        int `json:"reward"`
}

type LoadingTip struct {
	Code           int    `json:"code"`
	LoadingTipType string `json:"loadingTipType"`
	MinLv          int    `json:"minLv"`
	MaxLv          int    `json:"maxLv"`
	TextKey        string `json:"textKey"`
	ImageName      string `json:"imageName"`
}

type MasteryExp struct {
	Code           int    `json:"code"`
	ModeType       int    `json:"modeType"`
	ConditionType  string `json:"conditionType"`
	Grade          string `json:"grade"`
	ConditionValue int    `json:"conditionValue"`
	MasteryType1   string `json:"masteryType1"`
	Value1         int    `json:"value1"`
	MasteryType2   string `json:"masteryType2"`
	Value2         int    `json:"value2"`
	MasteryType3   string `json:"masteryType3"`
	Value3         int    `json:"value3"`
}

type MasteryLevel struct {
	Code              int    `json:"code"`
	Type              string `json:"type"`
	MasteryLevel      int    `json:"masteryLevel"`
	NextMasteryExp    int    `json:"nextMasteryExp"`
	GiveLevelExp      int    `json:"giveLevelExp"`
	WeaponSkillPoint  int    `json:"weaponSkillPoint"`
	ExpGrowthCapRatio int    `json:"expGrowthCapRatio"`
}

type MasteryStat struct {
	Code                      int     `json:"code"`
	Type                      string  `json:"type"`
	CharacterCode             int     `json:"characterCode"`
	FirstOption               string  `json:"firstOption"`
	FirstOptionSection1Value  float64 `json:"firstOptionSection1Value"`
	FirstOptionSection2Value  float64 `json:"firstOptionSection2Value"`
	FirstOptionSection3Value  float64 `json:"firstOptionSection3Value"`
	FirstOptionSection4Value  float64 `json:"firstOptionSection4Value"`
	SecondOption              string  `json:"secondOption"`
	SecondOptionSection1Value float64 `json:"secondOptionSection1Value"`
	SecondOptionSection2Value float64 `json:"secondOptionSection2Value"`
	SecondOptionSection3Value float64 `json:"secondOptionSection3Value"`
	SecondOptionSection4Value float64 `json:"secondOptionSection4Value"`
	ThirdOption               string  `json:"thirdOption"`
	ThirdOptionSection1Value  int     `json:"thirdOptionSection1Value"`
	ThirdOptionSection2Value  int     `json:"thirdOptionSection2Value"`
	ThirdOptionSection3Value  int     `json:"thirdOptionSection3Value"`
	ThirdOptionSection4Value  int     `json:"thirdOptionSection4Value"`
}

type Monster struct {
	Code             int     `json:"Code"`
	Monster          string  `json:"monster"`
	IsMutant         bool    `json:"isMutant"`
	Grade            string  `json:"grade"`
	Mode             int     `json:"mode"`
	CreateDay        string  `json:"createDay"`
	CreateTime       int     `json:"createTime"`
	RegenTime        int     `json:"regenTime"`
	LevelUpPeriod    int     `json:"levelUpPeriod"`
	LevelUpAmount    int     `json:"levelUpAmount"`
	LevelUpMax       int     `json:"levelUpMax"`
	MaxHp            int     `json:"maxHp"`
	MaxEp            int     `json:"maxEp"`
	InitExtraPoint   int     `json:"initExtraPoint"`
	AttackPower      int     `json:"attackPower"`
	Defense          int     `json:"defense"`
	AttackSpeed      float64 `json:"attackSpeed"`
	MoveSpeed        int     `json:"moveSpeed"`
	SightRange       int     `json:"sightRange"`
	ChasingRange     int     `json:"chasingRange"`
	AttackRange      float64 `json:"attackRange"`
	FirstAttackRange int     `json:"firstAttackRange"`
	Aggressive       string  `json:"aggressive"`
	DetectInvisible  bool    `json:"detectInvisible"`
	Radius           int     `json:"radius"`
	PathingRadius    float64 `json:"pathingRadius"`
	UIHeight         int     `json:"uiHeight"`
	GainExp          int     `json:"gainExp"`
	TargetOnRange    int     `json:"targetOnRange"`
	RandomDropCount  int     `json:"randomDropCount"`
	Resource         string  `json:"resource"`
	CorpseResource   string  `json:"corpseResource"`
	AppearTime       float64 `json:"appearTime"`
}

type MonsterDropGroup struct {
	MonsterCode  int `json:"monsterCode"`
	MonsterLevel int `json:"monsterLevel"`
	DropGroup    int `json:"dropGroup"`
}

type MonsterLevelUpStat struct {
	Code        int     `json:"code"`
	Monster     string  `json:"monster"`
	Mode        int     `json:"mode"`
	MaxHp       int     `json:"maxHp"`
	AttackPower int     `json:"attackPower"`
	Defense     float64 `json:"defense"`
	MoveSpeed   float64 `json:"moveSpeed"`
	GainExp     int     `json:"gainExp"`
}

type MonsterSpawnLevel struct {
	Code        int `json:"code"`
	Mode        int `json:"mode"`
	PlayerLevel int `json:"playerLevel"`
	MonsterCode int `json:"monsterCode"`
	SpawnLevel  int `json:"spawnLevel"`
}

type NaviCollectAndHunt struct {
	Code         int    `json:"code"`
	ItemCode     int    `json:"itemCode"`
	AreaCodeList string `json:"areaCodeList"`
}

type NearByArea struct {
	Code           int `json:"code"`
	AreaCode       int `json:"areaCode"`
	NearByAreaCode int `json:"nearByAreaCode"`
}

type RandomEquipment struct {
	Code          int    `json:"code"`
	Group         string `json:"group"`
	Itemcode      int    `json:"itemcode"`
	Weight        int    `json:"weight"`
	ItemGrade     string `json:"itemGrade"`
	TagMultiplier int    `json:"tagMultiplier"`
	CharacterNum  int    `json:"characterNum"`
}

type RecommendedList struct {
	Code                  int    `json:"code"`
	Character             string `json:"character"`
	CharacterCode         int    `json:"characterCode"`
	Mastery               string `json:"mastery"`
	StartWeapon           int    `json:"startWeapon"`
	CobaltStartWeapon     int    `json:"cobaltStartWeapon"`
	StartItemGroupCode    int    `json:"startItemGroupCode"`
	CobaltDraft           int    `json:"cobaltDraft"`
	CobaltExtraDraft      string `json:"cobaltExtraDraft"`
	CobaltCanChooseWeapon bool   `json:"cobaltCanChooseWeapon"`
	FavoriteMainTag       string `json:"favoriteMainTag"`
}

type Season struct {
	SeasonID    int    `json:"seasonID"`
	SeasonName  string `json:"seasonName"`
	SeasonStart string `json:"seasonStart"`
	SeasonEnd   string `json:"seasonEnd"`
	IsCurrent   int    `json:"isCurrent"`
}

type SummonObjectStat struct {
	Code                 int     `json:"code"`
	Name                 string  `json:"name"`
	Duration             int     `json:"duration"`
	CreateRange          int     `json:"createRange"`
	PileRange            int     `json:"pileRange"`
	CreateVisibleTime    int     `json:"createVisibleTime"`
	CreateStealthTime    int     `json:"createStealthTime"`
	InfiltrationTime     int     `json:"infiltrationTime"`
	DetectionRange       int     `json:"detectionRange"`
	MaxHp                int     `json:"maxHp"`
	MaxSp                int     `json:"maxSp"`
	InitExtraPoint       int     `json:"initExtraPoint"`
	MaxExtraPoint        int     `json:"maxExtraPoint"`
	AttackPower          int     `json:"attackPower"`
	Defense              int     `json:"defense"`
	CriticalStrikeChance int     `json:"criticalStrikeChance"`
	HpRegen              int     `json:"hpRegen"`
	SpRegen              int     `json:"spRegen"`
	RangeRadius          int     `json:"rangeRadius"`
	AttackSpeed          int     `json:"attackSpeed"`
	AttackRange          int     `json:"attackRange"`
	AttackDelay          int     `json:"attackDelay"`
	MoveSpeed            int     `json:"moveSpeed"`
	Radius               float64 `json:"radius"`
	UIHeight             float64 `json:"uiHeight"`
	SightRange           int     `json:"sightRange"`
	SightAngle           int     `json:"sightAngle"`
}

type TacticalSkillSet struct {
	Code            int `json:"code"`
	NextUpgradecode int `json:"nextUpgradecode"`
	UpgradeCredit   int `json:"upgradeCredit"`
	UpgradeMaterial int `json:"upgradeMaterial"`
	SkillCode       int `json:"skillCode"`
}

type TacticalSkillSetGroup struct {
	Group          int    `json:"group"`
	ModeType       int    `json:"modeType"`
	StartCode      int    `json:"startCode"`
	EquipWithStart bool   `json:"equipWithStart"`
	Icon           string `json:"icon"`
}

type Trait struct {
	Code          int    `json:"code"`
	OpenAccountLv int    `json:"openAccountLv"`
	TraitGroup    string `json:"traitGroup"`
	TraitType     string `json:"traitType"`
	Active        bool   `json:"active"`
}

type TransferConsole struct {
	ItemCode                          int    `json:"itemCode"`
	Mode                              int    `json:"mode"`
	ItemType                          string `json:"itemType"`
	TransferTimeSafeArea              int    `json:"transferTimeSafeArea"`
	SubtractTrasferTimeRestrictedArea int    `json:"subtractTrasferTimeRestrictedArea"`
	ManufactureCooldown               int    `json:"manufactureCooldown"`
	ConsumeVFCredit                   int    `json:"consumeVFCredit"`
	LimitCount                        int    `json:"limitCount"`
	TraitSale                         bool   `json:"traitSale"`
}

type VFCredit struct {
	Code           int    `json:"code"`
	Mode           int    `json:"mode"`
	Phase          int    `json:"phase"`
	ConditionType  string `json:"conditionType"`
	ConditionValue int    `json:"conditionValue"`
	AcquireSelf    int    `json:"acquireSelf"`
	AcquireTeam    int    `json:"acquireTeam"`
}

type WeaponTypeInfo struct {
	Type                  string  `json:"type"`
	AttackSpeed           float64 `json:"attackSpeed"`
	AttackRange           float64 `json:"attackRange"`
	ShopFilter            int     `json:"shopFilter"`
	SummonObjectHitDamage int     `json:"summonObjectHitDamage"`
}

type GameInfo struct {
	UserNum                   int         `json:"userNum"`
	Nickname                  string      `json:"nickname"`
	GameID                    int         `json:"gameId"`
	SeasonID                  int         `json:"seasonId"`
	MatchingMode              int         `json:"matchingMode"`
	MatchingTeamMode          int         `json:"matchingTeamMode"`
	CharacterNum              int         `json:"characterNum"`
	SkinCode                  int         `json:"skinCode"`
	CharacterLevel            int         `json:"characterLevel"`
	GameRank                  int         `json:"gameRank"`
	PlayerKill                int         `json:"playerKill"`
	PlayerAssistant           int         `json:"playerAssistant"`
	MonsterKill               int         `json:"monsterKill"`
	BestWeapon                int         `json:"bestWeapon"`
	BestWeaponLevel           int         `json:"bestWeaponLevel"`
	MasteryLevel              map[int]int `json:"masteryLevel"`
	Equipment                 map[int]int `json:"equipment"`
	VersionMajor              int         `json:"versionMajor"`
	VersionMinor              int         `json:"versionMinor"`
	Language                  string      `json:"language"`
	SkillLevelInfo            map[int]int `json:"skillLevelInfo"`
	SkillOrderInfo            map[int]int `json:"skillOrderInfo"`
	ServerName                string      `json:"serverName"`
	MaxHp                     int         `json:"maxHp"`
	MaxSp                     int         `json:"maxSp"`
	AttackPower               int         `json:"attackPower"`
	Defense                   int         `json:"defense"`
	HpRegen                   float64     `json:"hpRegen"`
	SpRegen                   float64     `json:"spRegen"`
	AttackSpeed               float64     `json:"attackSpeed"`
	MoveSpeed                 float64     `json:"moveSpeed"`
	OutOfCombatMoveSpeed      float64     `json:"outOfCombatMoveSpeed"`
	SightRange                float64     `json:"sightRange"`
	AttackRange               float64     `json:"attackRange"`
	CriticalStrikeChance      int         `json:"criticalStrikeChance"`
	CriticalStrikeDamage      int         `json:"criticalStrikeDamage"`
	CoolDownReduction         int         `json:"coolDownReduction"`
	LifeSteal                 float64     `json:"lifeSteal"`
	NormalLifeSteal           int         `json:"normalLifeSteal"`
	SkillLifeSteal            int         `json:"skillLifeSteal"`
	AmplifierToMonster        int         `json:"amplifierToMonster"`
	TrapDamage                int         `json:"trapDamage"`
	BonusCoin                 int         `json:"bonusCoin"`
	GainExp                   int         `json:"gainExp"`
	BaseExp                   int         `json:"baseExp"`
	BonusExp                  int         `json:"bonusExp"`
	StartDtm                  string      `json:"startDtm"`
	Duration                  int         `json:"duration"`
	PlayTime                  int         `json:"playTime"`
	WatchTime                 int         `json:"watchTime"`
	TotalTime                 int         `json:"totalTime"`
	SurvivableTime            int         `json:"survivableTime"`
	BotAdded                  int         `json:"botAdded"`
	BotRemain                 int         `json:"botRemain"`
	RestrictedAreaAccelerated int         `json:"restrictedAreaAccelerated"`
	SafeAreas                 int         `json:"safeAreas"`
	TeamNumber                int         `json:"teamNumber"`
	PreMade                   int         `json:"preMade"`
	EventMissionResult        struct {
	} `json:"eventMissionResult"`
	GainedNormalMmrKFactor      float64     `json:"gainedNormalMmrKFactor"`
	Victory                     int         `json:"victory"`
	CraftUncommon               int         `json:"craftUncommon"`
	CraftRare                   int         `json:"craftRare"`
	CraftEpic                   int         `json:"craftEpic"`
	CraftLegend                 int         `json:"craftLegend"`
	DamageToPlayer              int         `json:"damageToPlayer"`
	DamageToPlayerTrap          int         `json:"damageToPlayer_trap"`
	DamageToPlayerBasic         int         `json:"damageToPlayer_basic"`
	DamageToPlayerSkill         int         `json:"damageToPlayer_skill"`
	DamageToPlayerItemSkill     int         `json:"damageToPlayer_itemSkill"`
	DamageToPlayerDirect        int         `json:"damageToPlayer_direct"`
	DamageToPlayerUniqueSkill   int         `json:"damageToPlayer_uniqueSkill"`
	DamageFromPlayer            int         `json:"damageFromPlayer"`
	DamageFromPlayerTrap        int         `json:"damageFromPlayer_trap"`
	DamageFromPlayerBasic       int         `json:"damageFromPlayer_basic"`
	DamageFromPlayerSkill       int         `json:"damageFromPlayer_skill"`
	DamageFromPlayerItemSkill   int         `json:"damageFromPlayer_itemSkill"`
	DamageFromPlayerDirect      int         `json:"damageFromPlayer_direct"`
	DamageFromPlayerUniqueSkill int         `json:"damageFromPlayer_uniqueSkill"`
	DamageToMonster             int         `json:"damageToMonster"`
	DamageToMonsterTrap         int         `json:"damageToMonster_trap"`
	DamageToMonsterBasic        int         `json:"damageToMonster_basic"`
	DamageToMonsterSkill        int         `json:"damageToMonster_skill"`
	DamageToMonsterItemSkill    int         `json:"damageToMonster_itemSkill"`
	DamageToMonsterDirect       int         `json:"damageToMonster_direct"`
	DamageToMonsterUniqueSkill  int         `json:"damageToMonster_uniqueSkill"`
	DamageFromMonster           int         `json:"damageFromMonster"`
	KillMonsters                map[int]int `json:"killMonsters"`
	HealAmount                  int         `json:"healAmount"`
	TeamRecover                 int         `json:"teamRecover"`
	ProtectAbsorb               int         `json:"protectAbsorb"`
	AddSurveillanceCamera       int         `json:"addSurveillanceCamera"`
	AddTelephotoCamera          int         `json:"addTelephotoCamera"`
	RemoveSurveillanceCamera    int         `json:"removeSurveillanceCamera"`
	RemoveTelephotoCamera       int         `json:"removeTelephotoCamera"`
	UseHyperLoop                int         `json:"useHyperLoop"`
	UseSecurityConsole          int         `json:"useSecurityConsole"`
	GiveUp                      int         `json:"giveUp"`
	TeamSpectator               int         `json:"teamSpectator"`
	RouteIDOfStart              int         `json:"routeIdOfStart"`
	RouteSlotID                 int         `json:"routeSlotId"`
	PlaceOfStart                string      `json:"placeOfStart"`
	MatchSize                   int         `json:"matchSize"`
	TeamKill                    int         `json:"teamKill"`
	TotalFieldKill              int         `json:"totalFieldKill"`
	AccountLevel                int         `json:"accountLevel"`
	KillerUserNum               int         `json:"killerUserNum"`
	Killer                      string      `json:"killer"`
	KillDetail                  string      `json:"killDetail"`
	CauseOfDeath                string      `json:"causeOfDeath"`
	PlaceOfDeath                string      `json:"placeOfDeath"`
	KillerCharacter             string      `json:"killerCharacter"`
	KillerWeapon                string      `json:"killerWeapon"`
	KillerUserNum2              int         `json:"killerUserNum2"`
	KillerUserNum3              int         `json:"killerUserNum3"`
	FishingCount                int         `json:"fishingCount"`
	UseEmoticonCount            int         `json:"useEmoticonCount"`
	ExpireDtm                   string      `json:"expireDtm"`
	TraitFirstCore              int         `json:"traitFirstCore"`
	TraitFirstSub               []int       `json:"traitFirstSub"`
	TraitSecondSub              []int       `json:"traitSecondSub"`
	AirSupplyOpenCount          []int       `json:"airSupplyOpenCount"`
	FoodCraftCount              []int       `json:"foodCraftCount"`
	BeverageCraftCount          []int       `json:"beverageCraftCount"`
	RankPoint                   int         `json:"rankPoint"`
	TotalTurbineTakeOver        int         `json:"totalTurbineTakeOver"`
	UsedNormalHealPack          int         `json:"usedNormalHealPack"`
	UsedReinforcedHealPack      int         `json:"usedReinforcedHealPack"`
	UsedNormalShieldPack        int         `json:"usedNormalShieldPack"`
	UsedReinforceShieldPack     int         `json:"usedReinforceShieldPack"`
	TotalVFCredits              []int       `json:"totalVFCredits"`
	UsedVFCredits               []int       `json:"usedVFCredits"`
	SumTotalVFCredits           int         `json:"sumTotalVFCredits"`
	SumUsedVFCredits            int         `json:"sumUsedVFCredits"`
	CraftMythic                 int         `json:"craftMythic"`
	PlayerDeaths                int         `json:"playerDeaths"`
	KillGamma                   bool        `json:"killGamma"`
	ScoredPoint                 []int       `json:"scoredPoint"`
	KillDetails                 string      `json:"killDetails"`
	DeathDetails                string      `json:"deathDetails"`
	KillsPhaseOne               int         `json:"killsPhaseOne"`
	KillsPhaseTwo               int         `json:"killsPhaseTwo"`
	KillsPhaseThree             int         `json:"killsPhaseThree"`
	DeathsPhaseOne              int         `json:"deathsPhaseOne"`
	DeathsPhaseTwo              int         `json:"deathsPhaseTwo"`
	DeathsPhaseThree            int         `json:"deathsPhaseThree"`
	UsedPairLoop                int         `json:"usedPairLoop"`
	CcTimeToPlayer              int         `json:"ccTimeToPlayer"`
	CreditSource                struct {
		PhaseStart                          int     `json:"PhaseStart"`
		TimeElapsedCompensationByMiliSecond int     `json:"TimeElapsedCompensationByMiliSecond"`
		KillChicken                         int     `json:"KillChicken"`
		TacticalSkillUpgrade                int     `json:"TacticalSkillUpgrade"`
		TimeElapsedCreditBonusByMiliSecond  float64 `json:"TimeElapsedCreditBonusByMiliSecond"`
		KillBat                             int     `json:"KillBat"`
		KillBoar                            int     `json:"KillBoar"`
		TransferConsoleRemoteDroneMySelf    int     `json:"TransferConsoleRemoteDroneMySelf"`
		KillBear                            int     `json:"KillBear"`
		KillAssistDivideContribute          int     `json:"KillAssistDivideContribute"`
		KillPlayerMerge                     int     `json:"KillPlayerMerge"`
		KillMutantChicken                   int     `json:"KillMutantChicken"`
		TransferConsoleSpecialMaterial      int     `json:"TransferConsoleSpecialMaterial"`
		KillOmega                           int     `json:"KillOmega"`
		KillWolf                            int     `json:"KillWolf"`
		KillWildDog                         int     `json:"KillWildDog"`
		KillMutantBat                       int     `json:"KillMutantBat"`
		KillMutantBear                      int     `json:"KillMutantBear"`
		KillWickline                        int     `json:"KillWickline"`
		ItemBounty                          int     `json:"ItemBounty"`
	} `json:"creditSource"`
	BoughtInfusion                          string      `json:"boughtInfusion"`
	ItemTransferredConsole                  []int       `json:"itemTransferredConsole"`
	ItemTransferredDrone                    []int       `json:"itemTransferredDrone"`
	Afk                                     bool        `json:"afk"`
	EscapeState                             int         `json:"escapeState"`
	TotalDoubleKill                         int         `json:"totalDoubleKill"`
	TotalTripleKill                         int         `json:"totalTripleKill"`
	TotalQuadraKill                         int         `json:"totalQuadraKill"`
	TotalExtraKill                          int         `json:"totalExtraKill"`
	CollectItemForLog                       []int       `json:"collectItemForLog"`
	EquipFirstItemForLog                    map[int]int `json:"equipFirstItemForLog"`
	BattleZone1AreaCode                     int         `json:"battleZone1AreaCode"`
	BattleZone1BattleMark                   int         `json:"battleZone1BattleMark"`
	BattleZone1ItemCode                     []any       `json:"battleZone1ItemCode"`
	BattleZone2AreaCode                     int         `json:"battleZone2AreaCode"`
	BattleZone2BattleMark                   int         `json:"battleZone2BattleMark"`
	BattleZone2ItemCode                     []any       `json:"battleZone2ItemCode"`
	BattleZone3AreaCode                     int         `json:"battleZone3AreaCode"`
	BattleZone3BattleMark                   int         `json:"battleZone3BattleMark"`
	BattleZone3ItemCode                     []any       `json:"battleZone3ItemCode"`
	BattleZonePlayerKill                    int         `json:"battleZonePlayerKill"`
	BattleZoneDeaths                        int         `json:"battleZoneDeaths"`
	BattleZone1Winner                       int         `json:"battleZone1Winner"`
	BattleZone2Winner                       int         `json:"battleZone2Winner"`
	BattleZone3Winner                       int         `json:"battleZone3Winner"`
	BattleZone1BattleMarkCount              int         `json:"battleZone1BattleMarkCount"`
	BattleZone2BattleMarkCount              int         `json:"battleZone2BattleMarkCount"`
	BattleZone3BattleMarkCount              int         `json:"battleZone3BattleMarkCount"`
	TacticalSkillGroup                      int         `json:"tacticalSkillGroup"`
	TacticalSkillLevel                      int         `json:"tacticalSkillLevel"`
	TotalGainVFCredit                       int         `json:"totalGainVFCredit"`
	KillPlayerGainVFCredit                  int         `json:"killPlayerGainVFCredit"`
	KillChickenGainVFCredit                 int         `json:"killChickenGainVFCredit"`
	KillBoarGainVFCredit                    int         `json:"killBoarGainVFCredit"`
	KillWildDogGainVFCredit                 int         `json:"killWildDogGainVFCredit"`
	KillWolfGainVFCredit                    int         `json:"killWolfGainVFCredit"`
	KillBearGainVFCredit                    int         `json:"killBearGainVFCredit"`
	KillOmegaGainVFCredit                   int         `json:"killOmegaGainVFCredit"`
	KillBatGainVFCredit                     int         `json:"killBatGainVFCredit"`
	KillWicklineGainVFCredit                int         `json:"killWicklineGainVFCredit"`
	KillAlphaGainVFCredit                   int         `json:"killAlphaGainVFCredit"`
	KillItemBountyGainVFCredit              int         `json:"killItemBountyGainVFCredit"`
	KillDroneGainVFCredit                   int         `json:"killDroneGainVFCredit"`
	KillGammaGainVFCredit                   int         `json:"killGammaGainVFCredit"`
	KillTurretGainVFCredit                  int         `json:"killTurretGainVFCredit"`
	ItemShredderGainVFCredit                int         `json:"itemShredderGainVFCredit"`
	TotalUseVFCredit                        int         `json:"totalUseVFCredit"`
	RemoteDroneUseVFCreditMySelf            int         `json:"remoteDroneUseVFCreditMySelf"`
	RemoteDroneUseVFCreditAlly              int         `json:"remoteDroneUseVFCreditAlly"`
	TransferConsoleFromMaterialUseVFCredit  int         `json:"transferConsoleFromMaterialUseVFCredit"`
	TransferConsoleFromEscapeKeyUseVFCredit int         `json:"transferConsoleFromEscapeKeyUseVFCredit"`
	TransferConsoleFromRevivalUseVFCredit   int         `json:"transferConsoleFromRevivalUseVFCredit"`
	TacticalSkillUpgradeUseVFCredit         int         `json:"tacticalSkillUpgradeUseVFCredit"`
	InfusionReRollUseVFCredit               int         `json:"infusionReRollUseVFCredit"`
	InfusionTraitUseVFCredit                int         `json:"infusionTraitUseVFCredit"`
	InfusionRelicUseVFCredit                int         `json:"infusionRelicUseVFCredit"`
	InfusionStoreUseVFCredit                int         `json:"infusionStoreUseVFCredit"`
	TeamElimination                         int         `json:"teamElimination"`
	TeamDown                                int         `json:"teamDown"`
	TeamBattleZoneDown                      int         `json:"teamBattleZoneDown"`
	TeamRepeatDown                          int         `json:"teamRepeatDown"`
	AdaptiveForce                           int         `json:"adaptiveForce"`
	AdaptiveForceAttack                     int         `json:"adaptiveForceAttack"`
	AdaptiveForceAmplify                    int         `json:"adaptiveForceAmplify"`
	SkillAmp                                int         `json:"skillAmp"`
	CampFireCraftUncommon                   int         `json:"campFireCraftUncommon"`
	CampFireCraftRare                       int         `json:"campFireCraftRare"`
	CampFireCraftEpic                       int         `json:"campFireCraftEpic"`
	CampFireCraftLegendary                  int         `json:"campFireCraftLegendary"`
	CobaltRandomPickRemoveCharacter         int         `json:"cobaltRandomPickRemoveCharacter"`
	TacticalSkillUseCount                   int         `json:"tacticalSkillUseCount"`
	CreditRevivalCount                      int         `json:"creditRevivalCount"`
	CreditRevivedOthersCount                int         `json:"creditRevivedOthersCount"`
	TimeSpentInBriefingRoom                 int         `json:"timeSpentInBriefingRoom"`
	IsLeavingBeforeCreditRevivalTerminate   bool        `json:"IsLeavingBeforeCreditRevivalTerminate"`
	CrGetAnimal                             int         `json:"crGetAnimal"`
	CrGetMutant                             int         `json:"crGetMutant"`
	CrGetPhaseStart                         int         `json:"crGetPhaseStart"`
	CrGetKill                               int         `json:"crGetKill"`
	CrGetAssist                             int         `json:"crGetAssist"`
	CrGetTimeElapsed                        int         `json:"crGetTimeElapsed"`
	CrGetCreditBonus                        int         `json:"crGetCreditBonus"`
	CrUseRemoteDrone                        int         `json:"crUseRemoteDrone"`
	CrUseUpgradeTacticalSkill               int         `json:"crUseUpgradeTacticalSkill"`
	CrUseTreeOfLife                         int         `json:"crUseTreeOfLife"`
	CrUseMeteorite                          int         `json:"crUseMeteorite"`
	CrUseMythril                            int         `json:"crUseMythril"`
	CrUseForceCore                          int         `json:"crUseForceCore"`
	CrUseVFBloodSample                      int         `json:"crUseVFBloodSample"`
	CrUseActivationModule                   int         `json:"crUseActivationModule"`
	CrUseRootkit                            int         `json:"crUseRootkit"`
	MmrGainInGame                           int         `json:"mmrGainInGame"`
	MmrLossEntryCost                        int         `json:"mmrLossEntryCost"`
	IsLeavingBeforeCreditRevivalTerminate0  bool        `json:"isLeavingBeforeCreditRevivalTerminate"`
}

type TopRankedPlayers struct {
	UserNum     int    `json:"userNum"`
	Nickname    string `json:"nickname"`
	Rank        int    `json:"rank"`
	Mmr         int    `json:"mmr"`
	UserEmblems []any  `json:"userEmblems"`
}

type UserRankInfo struct {
	UserNum  int    `json:"userNum"`
	Mmr      int    `json:"mmr"`
	Nickname string `json:"nickname"`
	Rank     int    `json:"rank"`
}

type UserInfo struct {
	UserNum  int    `json:"userNum"`
	Nickname string `json:"nickname"`
}

type UserStats struct {
	SeasonID          int                  `json:"seasonId"`
	UserNum           int                  `json:"userNum"`
	MatchingMode      int                  `json:"matchingMode"`
	MatchingTeamMode  int                  `json:"matchingTeamMode"`
	Mmr               int                  `json:"mmr"`
	Nickname          string               `json:"nickname"`
	Rank              int                  `json:"rank"`
	RankSize          int                  `json:"rankSize"`
	TotalGames        int                  `json:"totalGames"`
	TotalWins         int                  `json:"totalWins"`
	TotalTeamKills    int                  `json:"totalTeamKills"`
	TotalDeaths       int                  `json:"totalDeaths"`
	EscapeCount       int                  `json:"escapeCount"`
	RankPercent       float64              `json:"rankPercent"`
	AverageRank       float64              `json:"averageRank"`
	AverageKills      float64              `json:"averageKills"`
	AverageAssistants float64              `json:"averageAssistants"`
	AverageHunts      float64              `json:"averageHunts"`
	Top1              float64              `json:"top1"`
	Top2              float64              `json:"top2"`
	Top3              float64              `json:"top3"`
	Top5              float64              `json:"top5"`
	Top7              float64              `json:"top7"`
	CharacterStats    []UserCharacterStats `json:"characterStats"`
}

type UserCharacterStats struct {
	CharacterCode int `json:"characterCode"`
	TotalGames    int `json:"totalGames"`
	Usages        int `json:"usages"`
	MaxKillings   int `json:"maxKillings"`
	Top3          int `json:"top3"`
	Wins          int `json:"wins"`
	Top3Rate      int `json:"top3Rate"`
	AverageRank   int `json:"averageRank"`
}

type ItemBuild struct {
	RecommendWeaponRoute struct {
		ID                     int     `json:"id"`
		Title                  string  `json:"title"`
		UserNum                int     `json:"userNum"`
		UserNickname           string  `json:"userNickname"`
		CharacterCode          int     `json:"characterCode"`
		SlotID                 int     `json:"slotId"`
		WeaponType             int     `json:"weaponType"`
		WeaponCodes            string  `json:"weaponCodes"`
		TacticalSkillGroupCode int     `json:"tacticalSkillGroupCode"`
		Paths                  string  `json:"paths"`
		Count                  int     `json:"count"`
		Version                string  `json:"version"`
		TeamMode               int     `json:"teamMode"`
		LanguageCode           string  `json:"languageCode"`
		RouteVersion           int     `json:"routeVersion"`
		Share                  bool    `json:"share"`
		UpdateDtm              int64   `json:"updateDtm"`
		Like                   int     `json:"v2Like"`
		WinRate                float64 `json:"v2WinRate"`
		SeasonID               int     `json:"v2SeasonId"`
		AccumulateLike         int     `json:"v2AccumulateLike"`
		AccumulateWinRate      float64 `json:"v2AccumulateWinRate"`
		AccumulateSeasonID     int     `json:"v2AccumulateSeasonId"`
	} `json:"recommendWeaponRoute"`
	RecommendWeaponRouteDesc struct {
		RecommendWeaponRouteID int `json:"recommendWeaponRouteId"`
	} `json:"recommendWeaponRouteDesc"`
}
