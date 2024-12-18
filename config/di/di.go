package di

func Initialize() controllers {
	repos := InitializeRepositories()
	adapters := InitializeAdapters(repos)
	services := InitializeServices(adapters)
	useCases := InitializeUseCases(services)

	return InitializeControllers(useCases)
}
