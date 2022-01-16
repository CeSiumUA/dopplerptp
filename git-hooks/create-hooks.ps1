cd ../
Copy-Item "git-hooks/pre-commit" -Destination ".git/hooks/pre-commit"
Copy-Item "git-hooks/pre-push" -Destination ".git/hooks/pre-push"