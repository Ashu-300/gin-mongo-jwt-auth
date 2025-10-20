pipeline {
    agent any   // Runs on any available Jenkins agent (node)

    environment {
        // Define any environment variables you need
        APP_NAME = 'my-app'
        BRANCH = "${env.GIT_BRANCH}"
    }

    stages {

        stage('Checkout') {
            steps {
                echo "Checking out source code..."
                checkout scm  // Automatically checks out the repository where Jenkinsfile resides
            }
        }

        stage('Build') {
            steps {
                echo "Building the project..."
                // Example: If Node.js project
                sh 'npm install'
                // For Python: sh 'pip install -r requirements.txt'
                // For Go: sh 'go build ./...'
            }
        }

        stage('Test') {
            steps {
                echo "Running tests..."
                // Example test commands
                sh 'npm test'
                // or sh 'pytest'
            }
        }

        stage('Package') {
            steps {
                echo "Packaging application..."
                // Example: Create Docker image
                sh 'docker build -t my-app:latest .'
            }
        }

        stage('Deploy') {
            steps {
                echo "Deploying application..."
                // Example: Run Docker container
                sh 'docker run -d -p 8080:8080 my-app:latest'
            }
        }
    }

    post {
        success {
            echo "✅ Build and deployment successful!"
        }
        failure {
            echo "❌ Build or test failed!"
        }
    }
}
