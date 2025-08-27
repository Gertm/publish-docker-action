# Publish Docker Action

Publish Docker Action builds, creates tags and pushes docker image to your docker registry.

## Usage

This simple example uses `Dockerfile` in your workspace to build image, attach the `latest`
tag and push to docker default registry (docker.io). Repository name is your GitHub repository
name by default.

```yaml
- uses: Gertm/publish-docker-action@master
  with:
    username: ${{ secrets.DOCKER_USERNAME }}
    password: ${{ secrets.DOCKER_PASSWORD }}
```

Use `file` and `path` arguments to set docker build file or build context if they are not in the default workspace.

### Set up registry and repository name

You can set docker registry with `registry` argument. Change docker repository name with `repository` argument.
For example:

```yaml
- uses: Gertm/publish-docker-action@master
  with:
    username: ${{ secrets.DOCKER_USERNAME }}
    password: ${{ secrets.DOCKER_PASSWORD }}
    registry: docker.pkg.github.com
    repository: Gertm/publish-docker-action
```

This will build and push the tag `docker.pkg.github.com/Gertm/publish-docker-action:latest`.

### Tags

#### Static Tag List

You can use static tag list by providing `tags` argument. Concat multiple tag names with commas.

```yaml
- uses: Gertm/publish-docker-action@master
  with:
    username: ${{ secrets.DOCKER_USERNAME }}
    password: ${{ secrets.DOCKER_PASSWORD }}
    registry: docker.pkg.github.com
    repository: Gertm/publish-docker-action
    tags: latest,newest,master
```

This example builds the image, creates three tags, and pushes all of them to the registry.

* `docker.pkg.github.com/Gertm/publish-docker-action:latest`
* `docker.pkg.github.com/Gertm/publish-docker-action:newest`
* `docker.pkg.github.com/Gertm/publish-docker-action:master`

#### Auto Tag

Set `with.auto_tag: true` to allow action generate docker image tags automatically.

```yaml
- uses: Gertm/publish-docker-action@master
  with:
    username: ${{ secrets.DOCKER_USERNAME }}
    password: ${{ secrets.DOCKER_PASSWORD }}
    registry: docker.pkg.github.com
    repository: Gertm/publish-docker-action
    auto_tag: true
```

Generated tags vary with `refs` types:

* branch: uses the branch name as docker tag name (`master` branch is renamed to `latest`).
* pull request: attaches a `pr-` prefix to branch name asdocker image tag. To allow pull request build, you must set `with.allow_pull_request` to `true`.
* tag: checks if the tag name is valid semantic version format (prefix `v` is allowed). If not, it uses git tag name as docker image tag directly. Else it generates three tags based on the version number, each followed with pre-release information.

Examples:

| Git | Docker Tag |
| --- | --- |
| branch `master` | `latest` |
| branch  `2019/09/28-new-feature` | `2019-09-28-new-feature` (`/` is replaced to `-`) |
| pull request `master` | `pr-master` |
| tag `1.0.0` | `1`, `1.0`, `1.0.0` |
| tag `v1.0.0` | `1`, `1.0`, `1.0.0` |
| tag `v1.0.0-rc1` | `1-rc1`, `1.0-rc1`, `1.0.0-rc1` |
| tag `20190921-actions` | `20190921-actions` (not semantic version) |

Auto tagging will override `with.tags` list.

Additionally, there's an output value `tag` you can use [in your next steps](https://help.github.com/en/actions/reference/contexts-and-expression-syntax-for-github-actions#steps-context).

```yaml
- id: build
  uses: Gertm/publish-docker-action@master
  with:
    username: ${{ secrets.DOCKER_USERNAME }}
    password: ${{ secrets.DOCKER_PASSWORD }}
    registry: docker.pkg.github.com
    repository: Gertm/publish-docker-action
    auto_tag: true

- id: deploy
  env:
    NEW_VERSION: ${{ steps.build.outputs.tag }}
  run: |
    docker pull $NEW_VERSION
```

### Cache

Provide `with.cache` argument to build from cache.

### Build Args

Use `with.build_args` to provide docker build-time variables. Multiple variables must be separated by comma. 

```yaml
- uses: Gertm/publish-docker-action@master
  with:
    username: ${{ secrets.DOCKER_USERNAME }}
    password: ${{ secrets.DOCKER_PASSWORD }}
    registry: docker.pkg.github.com
    repository: Gertm/publish-docker-action
    build_args: HTTP_PROXY=http://127.0.0.1,USER=nginx
```

### Target for Multi-Stage Builds

Provide `with.target` argument to set `--target` flag for docker build.

## Note

Please use the latest released version rather than master.
