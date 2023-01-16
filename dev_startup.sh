cd ./Backend
go run ./command_line/ > web_backend_logs.txt &
cd ../frontend
npm run dev > web_frontend_logs.txt