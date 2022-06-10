# Harness Argo CD #

## Upgrade Documentation

0. Connect upstream
1. Pull upstream tags
- `git fetch --all --tags`   
2. Create a branch from a version tag
- `git checkout tags/vM.M.P -b harness/vM.M.P`
3. Cherry-pick : [Harness Modifications](https://harness.atlassian.net/wiki/spaces/GIT/pages/21086830726/Maintain+ArgoCD+fork+for+GitOps+Agent)
4. make harness-gen 
5. Upgrade Service & Agent