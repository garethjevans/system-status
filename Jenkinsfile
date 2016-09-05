node('golang') {
	stage 'Build'
	env.GOPATH = pwd()
	sh 'go get github.com/Masterminds/glide'
    sh './bin/glide update'
}
