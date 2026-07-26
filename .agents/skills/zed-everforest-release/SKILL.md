---
name: zed-everforest-release
description: Releases the zed-everforest theme and submits its registry update to zed-industries/extensions. Use whenever publishing, tagging, creating a GitHub release, bumping Everforest's Zed extension version, or opening the upstream extensions PR.
---

# Release Zed Everforest

Release this repository first, then update the Zed extensions registry. Keep the source release and registry pull request separate so the registry always points to an immutable, tested release commit.

## Inputs

Collect these values before changing files:

- `VERSION`: semantic version without a `v`, such as `0.3.0`.
- `TAG`: `v${VERSION}`.
- `REGISTRY_BRANCH`: `everforest-${VERSION//./-}`.
- Release summary: two to four user-visible changes.

Set the shell variables once and reuse them throughout the release:

```bash
VERSION="0.3.0"
TAG="v${VERSION}"
REGISTRY_BRANCH="everforest-${VERSION//./-}"
```

Use these fixed repositories:

- Extension source: `https://github.com/albertsko/zed-everforest`.
- Registry fork: `git@github.com:albertsko/zed-extensions.git`.
- Registry upstream: `https://github.com/zed-industries/extensions.git`.

## Release Gates

Stop before publishing when any gate fails:

1. Confirm the source worktree contains only intended release changes.
2. Confirm `extension.toml` and the registry will use the same `VERSION`.
3. Confirm the extension has been tested in Zed as a development extension when the release changes visual behavior.
4. Confirm `LICENSE` remains an accepted Zed extension license.
5. Run local checks before committing and wait for source repository CI before tagging.
6. Never move, replace, or force-push a published release tag. Publish a follow-up version instead.

## Release The Source Repository

1. Inspect the source repository before committing:

```bash
gh auth status
git status --short --branch
git diff --check
git diff
git log --oneline -10
git tag --list --sort=-version:refname
gh release list --limit 20
```

2. Regenerate and validate the themes:

```bash
make generate
make ci
jq empty themes/everforest-regular.json themes/everforest-material.json themes/everforest-blur.json
git diff --check
```

3. Set `version = "VERSION"` in `extension.toml`.

4. Commit source changes and the generated themes together. Use a concise commit subject that describes the actual change.

5. Commit the version separately to match repository history:

```bash
git add extension.toml
git commit -m "chore: bump to ${VERSION}"
```

6. Recheck the final commit sequence and require a clean worktree:

```bash
git status --short --branch
git log --oneline --decorate -5
```

7. Push `main`, locate CI for the exact release commit, and wait for success:

```bash
RELEASE_SHA="$(git rev-parse HEAD)"
git push origin main
gh run list --workflow ci.yml --branch main --commit "${RELEASE_SHA}"
gh run watch "${RUN_ID}" --exit-status
```

Poll `gh run list` until GitHub creates the run, then assign its database ID to `RUN_ID`.

8. Create an annotated tag only after CI succeeds:

```bash
git tag -a "${TAG}" -m "Release ${TAG}"
git push origin "${TAG}"
```

9. Verify the remote tag resolves to `RELEASE_SHA`, not merely to an annotated tag object:

```bash
git rev-list -n 1 "${TAG}"
git ls-remote --tags origin "refs/tags/${TAG}" "refs/tags/${TAG}^{}"
```

10. Create a stable GitHub release with curated notes. Do not use generated notes when they would include unrelated history:

```bash
gh release create "${TAG}" --verify-tag --latest --title "${TAG}" --notes "${RELEASE_NOTES}"
```

11. Save the release URL and verify that the release is neither a draft nor a prerelease:

```bash
gh release view "${TAG}" --json url,tagName,name,isDraft,isPrerelease,publishedAt
```

## Update The Zed Registry

Follow [Zed's developing and publishing guide](https://zed.dev/docs/extensions/developing-extensions). In particular, keep the submodule URL on HTTPS, point it to a commit available on a public branch, match both manifest versions, preserve an accepted license, and run the registry sorter.

1. Clone the registry fork into a fresh temporary directory. Never reuse a checkout with unrelated changes:

```bash
REGISTRY_DIR="$(mktemp -d "${TMPDIR:-/tmp}/zed-extensions.XXXXXX")"
git clone git@github.com:albertsko/zed-extensions.git "${REGISTRY_DIR}"
```

Run the remaining registry commands from `REGISTRY_DIR`.

2. Add and fetch upstream, then fast-forward the fork's `main`:

```bash
git remote add upstream https://github.com/zed-industries/extensions.git
git fetch upstream main
git fetch origin main
git checkout main
git merge --ff-only upstream/main
git push origin main
```

3. Create `REGISTRY_BRANCH` from synchronized `main`:

```bash
git switch -c "${REGISTRY_BRANCH}"
```

4. Advance the Everforest submodule and update the registry version:

```bash
git submodule update --init --remote extensions/everforest
```

Set the `[everforest]` entry in `extensions.toml` to `version = "VERSION"`.

5. Verify all release identities agree:

- `extensions.toml` reports `VERSION`.
- `extensions/everforest/extension.toml` reports `VERSION`.
- The Everforest gitlink resolves to `RELEASE_SHA`.
- `origin/main` in the submodule contains `RELEASE_SHA`.
- `.gitmodules` still uses `https://github.com/albertsko/zed-everforest.git`.

6. Install dependencies, sort registry files, and run local checks:

```bash
pnpm install --frozen-lockfile
pnpm sort-extensions
pnpm build
pnpm test
git diff --exit-code -- .gitmodules
git diff --check
```

7. Inspect status, diff, and recent history. Stage only `extensions.toml` and the Everforest gitlink:

```bash
git status --short --branch
git diff --submodule=log -- extensions.toml extensions/everforest
git log --oneline -10
git add extensions.toml extensions/everforest
git commit -m "Update everforest to v${VERSION}"
```

8. Push the branch and perform the final pull request review against the latest upstream `main`:

```bash
git push --set-upstream origin "${REGISTRY_BRANCH}"
git fetch upstream main
git status --short --branch
git rev-list --left-right --count upstream/main...HEAD
git diff --check upstream/main...HEAD
git diff --submodule=log upstream/main...HEAD -- extensions.toml extensions/everforest
```

9. Open a simple upstream pull request. Keep the body to the release URL only:

```bash
gh pr create --repo zed-industries/extensions --base main --head "albertsko:${REGISTRY_BRANCH}" --title "Update everforest to ${VERSION}" --body "${RELEASE_URL}"
```

10. Verify the title, body, files, base, head, and checks with `gh pr view` and `gh pr checks`.

## Done Criteria

Report these outcomes:

- Source release URL.
- Source CI URL and conclusion.
- Release commit and annotated tag.
- Registry commit and branch.
- Upstream pull request URL.
- Local validation results and any remaining pull request checks.

## Example

For `VERSION=0.3.0`, derive:

```text
TAG=v0.3.0
REGISTRY_BRANCH=everforest-0-3-0
PR title: Update everforest to 0.3.0
PR body: https://github.com/albertsko/zed-everforest/releases/tag/v0.3.0
```
