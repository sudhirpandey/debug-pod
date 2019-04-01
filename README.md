
### Hello world debug app

1) Build the binary form the code (static compile)
	```
	go build -tags netgo
	```

2) build the app
	```
	docker build -t docker-registry-default.svc/debug-proj/go-hello .
	```

3) Push the image into regsitry (internal registry)
	```
	docker push docker-registry-default.svc/debug-proj/go-hello
	```

4) Deploy the app from the image steam created form push above
	```
	oc new-app --image-stream=debug-proj/go-hello
	```

5) Validate the pods get depployed 
	```
	oc logs go-hello-xxx 
	```
