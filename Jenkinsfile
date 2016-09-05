node('golang') {
	env.GOPATH = pwd()
    stage 'Prereq'
    sh 'rm -fr *'
    git 'https://github.com/garethjevans/system-status'
	sh 'go get github.com/tebeka/go2xunit'
	sh 'go get -d'

	stage 'Build'
	sh 'go build'

    stage 'Test'
	sh 'go test -v'
    sh 'go test -v > test.output'

    stage 'Report'
    sh 'cat test.output | ./bin/go2xunit -output tests.xml'
	step([$class: 'JUnitResultArchiver', testResults: 'tests.xml'])
}

