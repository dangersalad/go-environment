package environment

var (
	// EnvKeyDeployEnv is the env variable name for the deployment environment
	EnvKeyDeployEnv = "DEPLOY_ENV"
	// DeployEnvDevelopment is the value for development deployments
	DeployEnvDevelopment = "development"
	// DeployEnvStaging is the value for staging deployments
	DeployEnvStaging = "staging"
	// DeployEnvProduction is the value for production deployments
	DeployEnvProduction = "production"
)

var currentDeployEnv string

// DeployEnv returns the current deploy env
func DeployEnv() string {
	if currentDeployEnv != "" {
		return currentDeployEnv
	}
	e, err := ReadOptions(Options{
		EnvKeyDeployEnv: "",
	})
	if err != nil {
		panic(err)
	}

	currentDeployEnv = e[EnvKeyDeployEnv]
	return currentDeployEnv
}
