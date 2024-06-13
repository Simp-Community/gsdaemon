package constants

// Constants related to plugin for GSTools
const (
	// the directory inside GSTools working directory that stores plugins' configuration files
	PLUGIN_CONFIG_DIRECTORY = "config"

	// The file prefix for a solo plugin (a single .py file)
	SOLO_PLUGIN_FILE_SUFFIX = ".go"

	PACKED_PLUGIN_FILE_SUFFIX1 = ".gst"
	PACKED_PLUGIN_FILE_SUFFIX2 = ".goz"

	// The name of the meta file inside a packed plugin
	PLUGIN_META_FILE = "gstools.plugin.json"

	// The name of the link file for a LinkedDirectoryPlugin
	LINK_DIRECTORY_PLUGIN_FILE_NAME = "gstools.linked_directory_plugin.json"

	// The name of the python package requirement file
	PLUGIN_REQUIREMENTS_FILE = "requirements.txt"

	// the path to the directory inside a packed plugin as the default directory to store the language files of the plugin
	PLUGIN_TRANSLATION_FILES_PATH = "lang"

	// Disabled plugin file name suffix
	DISABLED_PLUGIN_FILE_SUFFIX = ".disabled"
)
