node('golang') {
    stage 'Prereq'
    sh 'rm -fr *'
    git 'https://github.com/garethjevans/system-status'

	stage 'Build'
	env.GOPATH = pwd()
	//sh 'go get github.com/Masterminds/glide'
    //sh 'tree'
    //sh './bin/glide update'
	sh 'go get -d'
	sh 'go build -x'

    stage 'Test'
    sh 'go test -v ./...'
}

