import docker

client = docker.from_env()
print("Starting execution")
client.containers.run('busybox', tty=True, remove=True, stdin_open=True)
print("Started Container")