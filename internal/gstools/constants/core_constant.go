package constants

const (
	NAME_SHORT   = "GST"
	NAME         = "GSTools"
	PACKAGE_NAME = "gstools"
	CLI_COMMAND  = PACKAGE_NAME

	VERSION = "1.0.0" // semver (MAJOR.MINOR.PATCH)

	GITHUB_URL        = "https://github.com/Simp-Community/GSTools"
	GITHUB_API_LATEST = "https://api.github.com/repos/Simp-Community/GSTools/releases/latest"
	DOCUMENTATION_URL = "https://gstools.readthedocs.io"

	PACKAGE_PATH         = ""
	LOGGING_FILE         = "logs/GSTools.log"
	LANGUAGE_FILE_SUFFIX = ".yml"
	DEFAULT_LANGUAGE     = "en_us"

	PLUGIN_THREAD_POOL_SIZE               = 4
	MAX_TASK_QUEUE_SIZE                   = 2048
	WAIT_TIME_AFTER_SERVER_STDOUT_END_SEC = 60
	REACTOR_QUEUE_FULL_WARN_INTERVAL_SEC  = 5
)
