package er_api

// Routes
const BASE_ROUTE = "https://open-api.bser.io/v1/"                                    //
const DATA_ROUTE = "https://open-api.bser.io/v2/data/"                               // https://open-api.bser.io/v2/data/{metaType}
const FREE_CHAR_ROUTE = "https://open-api.bser.io/v1/freeCharacters/2"               //
const LANGUAGE_ROUTE = "https://open-api.bser.io/v1/l10n/"                           // https://open-api.bser.io/v1/l10n/{language}
const GAME_ROUTE = "https://open-api.bser.io/v1/games/"                              // https://open-api.bser.io/v1/games/{gameId}
const RANK_TOP_ROUTE = "https://open-api.bser.io/v1/rank/top/"                       // https://open-api.bser.io/v1/rank/top/{seasonId}/{matchingTeamMode}
const RANK_USER_ROUTE = "https://open-api.bser.io/v1/rank/"                          // https://open-api.bser.io/v1/rank/{userNum}/{seasonId}/{matchingTeamMode}
const USER_GAMES_ROUTE = "https://open-api.bser.io/v1/user/games/"                   // https://open-api.bser.io/v1/user/games/{userNum}
const USER_NICKNAME_ROUTE = "https://open-api.bser.io/v1/user/nickname"              // ?query={nickname}
const USER_STATS = "https://open-api.bser.io/v1/user/stats/"                         // https://open-api.bser.io/v1/user/stats/{userNum}/{seasonId}
const RECOMMENDED_BUILD_ROUTE = "https://open-api.bser.io/v1/weaponRoutes/recommend" // ?query={nextId}
const FIND_BUILD_ROUTE = "https://open-api.bser.io/v1/weaponRoutes/recommend/"       // https://open-api.bser.io/v1/weaponRoutes/recommend/{routeId}

// Language Options
const ENGLISH = "English"
const KOREAN = "Korean"
const JAPANESE = "Japanese"
const CHINESE_SIMPLIFIED = "ChineseSimplified"
const CHINESE_TRADITIONAL = "ChineseTraditional"
const FRENCH = "French"
const SPANISH = "Spanish"
const SPANISH_LATIN = "SpanishLatin"
const PORTUGUESE = "Portuguese"
const PORTUGUESE_LATIN = "PortugueseLatin"
const INDONESIAN = "Indonesian"
const GERMAN = "German"
const RUSSIAN = "Russian"
const THAI = "Thai"
const VIETNAMESE = "Vietnamese"

func getLanguages() [15]string {
	return [15]string{
		ENGLISH,
		KOREAN,
		JAPANESE,
		CHINESE_SIMPLIFIED,
		CHINESE_TRADITIONAL,
		FRENCH,
		SPANISH,
		SPANISH_LATIN,
		PORTUGUESE,
		PORTUGUESE_LATIN,
		INDONESIAN,
		GERMAN,
		RUSSIAN,
		THAI,
		VIETNAMESE,
	}
}

// Data V2 Meta Types
const DATA_ACTION_COST string = "ActionCost"
const DATA_AREA string = "Area"
const DATA_BATTLE_ZONE_REWARDS string = "BattleZoneReward"
const DATA_BULLET_CAPACITY string = "BulletCapacity"
const DATA_CHARACTER string = "Character"
const DATA_CHARACTER_ATTRIBUTES string = "CharacterAttributes"
const DATA_CHARACTER_EXPERIENCE string = "CharacterExp"
const DATA_CHARACTER_LEVEL_UP_STAT string = "CharacterLevelUpStat"
const DATA_CHARACTER_MASTERY string = "CharacterMastery"
const DATA_CHARACTER_MODE_MODIFIER string = "CharacterModeModifier"
const DATA_CHARACTER_SKIN string = "CharacterSkin"
const DATA_COLLECTIBLE string = "Collectible"
const DATA_DROP_GROUP string = "DropGroup"
const DATA_GAIN_EXP string = "GainExp"
const DATA_GAIN_SCORE string = "GainScore"
const DATA_GAME_TIP string = "GameTip"
const DATA_HOW_TO_FIND_ITEM string = "HowToFindItem"
const DATA_INFUSION_PRODUCT string = "InfusionProduct"
const DATA_ITEM_ARMOR string = "ItemArmor"
const DATA_ITEM_CONSUMABLE string = "ItemConsumable"
const DATA_ITEM_MISC string = "ItemMisc"
const DATA_ITEM_SEARCH_OPTION string = "ItemSearchOptionV2"
const DATA_ITEM_SPAWN string = "ItemSpawn"
const DATA_ITEM_SPECIAL string = "ItemSpecial"
const DATA_ITEM_WEAPON string = "ItemWeapon"
const DATA_ACCOUNT_LEVEL string = "Level"
const DATA_LOADING_TIP string = "LoadingTip"
const DATA_MASTERY_EXP string = "MasteryExp"
const DATA_MASTERY_LEVEL string = "MasteryLevel"
const DATA_MASTERY_STAT string = "MasteryStat"
const DATA_MONSTER string = "Monster"
const DATA_MONSTER_DROP_GROUP string = "MonsterDropGroup"
const DATA_MONSTER_LEVEL_UP_STAT string = "MonsterLevelUpStat"
const DATA_MONSTER_SPAWN_LEVEL string = "MonsterSpawnLevel"
const DATA_NAVI_COLLECT_AND_HUNT string = "NaviCollectAndHunt"
const DATA_NEAR_BY_AREA string = "NearByArea"
const DATA_RANDOM_EQUIPMENT string = "RandomEquipment"
const DATA_RECOMMENDED_LIST string = "RecommendedList"
const DATA_SEASONS string = "Season"
const DATA_SUMMON_OBJECT_STAT string = "SummonObjectStat"
const DATA_TACTICAL_SKILL_SET string = "TacticalSkillSet"
const DATA_TACTICAL_SKILL_SET_GROUP string = "TacticalSkillSetGroup"
const DATA_AUGMENTS string = "Trait"
const DATA_TRANSFER_CONSOLE string = "TransferConsole"
const DATA_VF_CREDIT string = "VFCredit"
const DATA_WEAPON_TYPE_INFO string = "WeaponTypeInfo"

/* Data V1 Meta Types - Probably do not use these but I'm not ready to remove them for real because I
realized too late that there's a V2 api so...........
const DATA_AREA_ATTRIBUTES string = "AreaAttributes"
const DATA_CHARACTER_SKILL_VIDEOS string = "CharacterSkillVideos"
const DATA_CHARACTER_STATE string = "CharacterState"
const DATA_CHARACTER_STATE_GROUP string = "CharacterStateGroup"
const DATA_CONTENT_OPEN string = "ContentOpen"
const DATA_CRITICAL_CHANGE string = "CriticalChance"
const DATA_EMBLEM string = "Emblem"
const DATA_EMOTES string = "Emotion"
const DATA_GAME_CONSTANT string = "GameConstant"
const DATA_GOODS string = "Goods"
const DATA_HOOK_LINE_PROJECTILE string = "HookLineProjectile"
const DATA_ITEM_SKILL string = "ItemSkill"
const DATA_ITEM_SKILL_GROUP string = "ItemSkillGroup"
const DATA_ITEM_SKILL_LINKER string = "ItemSkillLinker"
const DATA_LOWEST_RANK_ADJUST string = "LowestRankAdjust"
const DATA_MATCHING_QUEUE_TIER string = "MatchingQueueTier"
const DATA_MISSION string = "Mission"
const DATA_MUTANT_MONSTER_SPAWN string = "MutantMonsterSpawn"
const DATA_NOISE string = "Noise"
const DATA_PHASE string = "Phase"
const DATA_PREMADE_TEAM_TIER_RESTRICT string = "PreMadeTeamTierRestrict"
const DATA_PROJECTILE_DEFLECTOR_SETTING string = "ProjectileDeflectorSetting"
const DATA_PROJECTILE_SETTING string = "ProjectileSetting"
const DATA_RECOMMENDED_AREA string = "RecommendedArea"
const DATA_RECOMMENDED_ITEM_LIST string = "RecommendedItemList"
const DATA_RESTRICTED_AREA string = "RestrictedArea"
const DATA_SERVER_REGION string = "ServerRegion"
const DATA_SHOP_PRODUCT_ITEM string = "ShopProductItem"
const DATA_SKILL string = "Skill"
const DATA_SKILL_EVOLUTION string = "SkillEvolution"
const DATA_SKILL_EVOLUTION_POINT string = "SkillEvolutionPoint"
const DATA_SKILL_GROUP string = "SkillGroup"
const DATA_START_ITEM string = "StartItem"
const DATA_SUMMON_OBJECT string = "SummonObject"
const DATA_SUMMON_OBJECT_GROUP string = "SummonObjectGroup"
const DATA_WEAPON_ROUTE string = "WeaponRoute"
*/
