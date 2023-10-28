package er_api

func BuildERApiManager(apiKey string) ERApiManager {
	return ERApiManager{apiKey: apiKey}
}
