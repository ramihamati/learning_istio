start:
	@echo "Starting APIs..."
	@go run /mnt/extension/learning_istio/weather-api/main.go | awk '{print "\033[32m" $$0 "\033[0m"}' &
	@go run /mnt/extension/learning_istio/proxy-api/main.go | awk '{print "\033[34m" $$0 "\033[0m"}'&
	@wait

stop:
	@echo "Stopping all running Go processes..."
	@pkill -f "go run"

restart: stop start
