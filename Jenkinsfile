node('golang') {
    stage 'Prereq'
    sh 'rm -fr *'
    git 'https://github.com/garethjevans/system-status'
	sh 'go get -d'

	stage 'Build'
	env.GOPATH = pwd()
	sh 'go build -x'

    stage 'Test'
    sh 'go test -v > test.output'
}

