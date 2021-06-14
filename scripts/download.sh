echo "version $1"
mkdir -p "dist"
# apps
for app in 'gateway' 'message' 'subscribe' 'web' 'middle' 'spider' 'cron' 'workflow' 'task'
	do
	  echo "\n downloading... $app\n"
		curl -L https://github.com/tsundata/assistant/releases/download/v$1/$app-linux-amd64 --output ./dist/$app-linux-amd64
		chmod +x ./dist/$app-linux-amd64
  done
# agents
for agent in 'server' 'redis'
	do
	  echo "\n downloading... $agent\n"
		curl -L https://github.com/tsundata/assistant/releases/download/v$1/$agent-agent-linux-amd64 --output ./dist/$agent-agent-linux-amd64
		chmod +x ./dist/$agent-agent-linux-amd64
	done
