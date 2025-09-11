



<!DOCTYPE html>
<html class="gl-light ui-indigo with-top-bar with-header " lang="en">
<head prefix="og: http://ogp.me/ns#">
<meta charset="utf-8">
<meta content="IE=edge" http-equiv="X-UA-Compatible">
<meta content="width=device-width, initial-scale=1" name="viewport">
<title>pkg/basicauth/options.go · 1e27f1f9b9852379064db1ed41acae54a0b7c495 · OLCNE / prometheus / thanos · GitLab</title>
<script>
//<![CDATA[
window.gon={};gon.math_rendering_limits_enabled=true;gon.features={"inlineBlame":false,"blobOverflowMenu":false,"filterBlobPath":false,"blobRepositoryVueHeaderApp":false,"directoryCodeDropdownUpdates":false};
//]]>
</script>


<script>
//<![CDATA[
var gl = window.gl || {};
gl.startup_calls = null;
gl.startup_graphql_calls = [{"query":"query getBlobInfo(\n  $projectPath: ID!\n  $filePath: [String!]!\n  $ref: String!\n  $refType: RefType\n  $shouldFetchRawText: Boolean!\n) {\n  project(fullPath: $projectPath) {\n    __typename\n    id\n    repository {\n      __typename\n      empty\n      blobs(paths: $filePath, ref: $ref, refType: $refType) {\n        __typename\n        nodes {\n          __typename\n          id\n          webPath\n          name\n          size\n          rawSize\n          rawTextBlob @include(if: $shouldFetchRawText)\n          fileType\n          language\n          path\n          blamePath\n          editBlobPath\n          gitpodBlobUrl\n          ideEditPath\n          forkAndEditPath\n          ideForkAndEditPath\n          codeNavigationPath\n          projectBlobPathRoot\n          forkAndViewPath\n          environmentFormattedExternalUrl\n          environmentExternalUrlForRouteMap\n          canModifyBlob\n          canModifyBlobWithWebIde\n          canCurrentUserPushToBranch\n          archived\n          storedExternally\n          externalStorage\n          externalStorageUrl\n          rawPath\n          replacePath\n          pipelineEditorPath\n          simpleViewer {\n            fileType\n            tooLarge\n            type\n            renderError\n          }\n          richViewer {\n            fileType\n            tooLarge\n            type\n            renderError\n          }\n        }\n      }\n    }\n  }\n}\n","variables":{"projectPath":"OLCNE/Prometheus/thanos","ref":"1e27f1f9b9852379064db1ed41acae54a0b7c495","refType":null,"filePath":"pkg/basicauth/options.go","shouldFetchRawText":true}}];

if (gl.startup_calls && window.fetch) {
  Object.keys(gl.startup_calls).forEach(apiCall => {
   gl.startup_calls[apiCall] = {
      fetchCall: fetch(apiCall, {
        // Emulate XHR for Rails AJAX request checks
        headers: {
          'X-Requested-With': 'XMLHttpRequest'
        },
        // fetch won’t send cookies in older browsers, unless you set the credentials init option.
        // We set to `same-origin` which is default value in modern browsers.
        // See https://github.com/whatwg/fetch/pull/585 for more information.
        credentials: 'same-origin'
      })
    };
  });
}
if (gl.startup_graphql_calls && window.fetch) {
  const headers = {"X-CSRF-Token":"Tk9bnNz6uM7EjYyaqA5tDDfp8j56w9s2TbnO8DAzCZIh-oTXIRNl9gz39BMHY_megyyrJucnp9O4zSMgjcRF2Q","x-gitlab-feature-category":"source_code_management"};
  const url = `https://linux-git.oraclecorp.com/api/graphql`

  const opts = {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      ...headers,
    }
  };

  gl.startup_graphql_calls = gl.startup_graphql_calls.map(call => ({
    ...call,
    fetchCall: fetch(url, {
      ...opts,
      credentials: 'same-origin',
      body: JSON.stringify(call)
    })
  }))
}


//]]>
</script>

<link rel="prefetch" href="/assets/webpack/monaco.d13f99c0.chunk.js">

<link rel="stylesheet" href="/assets/application-f3a46613d0c62857b36ab461ec6c150580272b96dd876596e93bbc5cf6371923.css" />
<link rel="stylesheet" href="/assets/page_bundles/tree-36265195585af5e9c554d953e140f9c837af2e2eecb061d3f93d2857768eb286.css" /><link rel="stylesheet" href="/assets/page_bundles/projects-ac447098e0a84696db1e979fab09cb627a49e7c1fc06db30ac3b2b35acf07322.css" /><link rel="stylesheet" href="/assets/page_bundles/commit_description-1e2cba4dda3c7b30dd84924809020c569f1308dea51520fe1dd5d4ce31403195.css" /><link rel="stylesheet" href="/assets/page_bundles/work_items-f8b0433a074a1464d564be96c61f98abba72c32c594f4624f237b4e41b1a5679.css" /><link rel="stylesheet" href="/assets/page_bundles/notes_shared-8b38073ba71abd72b511b161ac1c0567ac634b689428f7c74591325093229970.css" />
<link rel="stylesheet" href="/assets/application_utilities-7844decc020ce5472253adae1ee2b329cdb6ddd08043b91daa3c9ed5530f5142.css" />
<link rel="stylesheet" href="/assets/tailwind-18b323869d1bc4b21818cb0e7f501054cab3084c92dc5d4f272982507e17363f.css" />


<link rel="stylesheet" href="/assets/fonts-fae5d3f79948bd85f18b6513a025f863b19636e85b09a1492907eb4b1bb0557b.css" />
<link rel="stylesheet" href="/assets/highlight/themes/white-99cce4f4b362f6840d7134d4129668929fde49c4da11d6ebf17f99768adbd868.css" />


<link rel="preload" href="/assets/application_utilities-7844decc020ce5472253adae1ee2b329cdb6ddd08043b91daa3c9ed5530f5142.css" as="style" type="text/css">
<link rel="preload" href="/assets/application-f3a46613d0c62857b36ab461ec6c150580272b96dd876596e93bbc5cf6371923.css" as="style" type="text/css">
<link rel="preload" href="/assets/highlight/themes/white-99cce4f4b362f6840d7134d4129668929fde49c4da11d6ebf17f99768adbd868.css" as="style" type="text/css">




<script src="/assets/webpack/runtime.c217ba1f.bundle.js" defer="defer"></script>
<script src="/assets/webpack/main.aa4dbbdf.chunk.js" defer="defer"></script>
<script src="/assets/webpack/commons-pages.groups.new-pages.import.gitlab_projects.new-pages.import.manifest.new-pages.projects.n-a0973272.dc791a28.chunk.js" defer="defer"></script>
<script src="/assets/webpack/commons-pages.groups.new-pages.import.gitlab_projects.new-pages.import.manifest.new-pages.projects.n-44c6c18e.5694c6d6.chunk.js" defer="defer"></script>
<script src="/assets/webpack/commons-pages.search.show-super_sidebar.09c5c921.chunk.js" defer="defer"></script>
<script src="/assets/webpack/super_sidebar.b794a9c8.chunk.js" defer="defer"></script>
<script src="/assets/webpack/commons-pages.projects-pages.projects.activity-pages.projects.alert_management.details-pages.project-2e472f70.fbe5726a.chunk.js" defer="defer"></script>
<script src="/assets/webpack/commons-pages.projects.blob.show-pages.projects.branches.new-pages.projects.commits.show-pages.proje-81161c0b.041430ca.chunk.js" defer="defer"></script>
<script src="/assets/webpack/commons-pages.import.bitbucket_server.new-pages.import.gitea.new-pages.import.gitlab_projects.new-pa-2584d496.5b38c1e7.chunk.js" defer="defer"></script>
<script src="/assets/webpack/commons-pages.projects.blob.show-pages.projects.show-pages.projects.snippets.edit-pages.projects.sni-42df7d4c.e2ff6bbf.chunk.js" defer="defer"></script>
<script src="/assets/webpack/72.a8ccd7b6.chunk.js" defer="defer"></script>
<script src="/assets/webpack/commons-pages.projects.blob.show-pages.projects.branches.index-pages.projects.show-pages.projects.ta-c9380b86.2bbbb6c7.chunk.js" defer="defer"></script>
<script src="/assets/webpack/commons-pages.projects.blob.show-pages.projects.show-pages.projects.snippets.show-pages.projects.tre-c684fcf6.3b3728c4.chunk.js" defer="defer"></script>
<script src="/assets/webpack/commons-pages.projects.blob.edit-pages.projects.blob.new-pages.projects.blob.show-pages.projects.sho-ec79e51c.c7cf396b.chunk.js" defer="defer"></script>
<script src="/assets/webpack/commons-pages.projects.blob.show-pages.projects.show-pages.projects.tree.show-pages.search.show.a57d5ac6.chunk.js" defer="defer"></script>
<script src="/assets/webpack/commons-pages.groups.show-pages.projects.blob.show-pages.projects.show-pages.projects.tree.show.44296e11.chunk.js" defer="defer"></script>
<script src="/assets/webpack/commons-pages.projects.blame.show-pages.projects.blob.show-pages.projects.show-pages.projects.tree.show.6fa9cae1.chunk.js" defer="defer"></script>
<script src="/assets/webpack/commons-pages.projects.blob.show-pages.projects.show-pages.projects.tree.show.214eb3ae.chunk.js" defer="defer"></script>
<script src="/assets/webpack/commons-pages.projects.blob.show-pages.projects.tree.show-treeList.d0c9b038.chunk.js" defer="defer"></script>
<script src="/assets/webpack/pages.projects.blob.show.6319b632.chunk.js" defer="defer"></script>

<meta content="object" property="og:type">
<meta content="GitLab" property="og:site_name">
<meta content="pkg/basicauth/options.go · 1e27f1f9b9852379064db1ed41acae54a0b7c495 · OLCNE / prometheus / thanos · GitLab" property="og:title">
<meta content="Thanos is a set of components that can be composed into a highly available metric system with unlimited storage capacity, which can be added seamlessly on top of..." property="og:description">
<meta content="https://linux-git.oraclecorp.com/assets/twitter_card-570ddb06edf56a2312253c5872489847a0f385112ddbcd71ccfa1570febab5d2.jpg" property="og:image">
<meta content="64" property="og:image:width">
<meta content="64" property="og:image:height">
<meta content="https://linux-git.oraclecorp.com/OLCNE/Prometheus/thanos/-/blob/1e27f1f9b9852379064db1ed41acae54a0b7c495/pkg/basicauth/options.go" property="og:url">
<meta content="summary" property="twitter:card">
<meta content="pkg/basicauth/options.go · 1e27f1f9b9852379064db1ed41acae54a0b7c495 · OLCNE / prometheus / thanos · GitLab" property="twitter:title">
<meta content="Thanos is a set of components that can be composed into a highly available metric system with unlimited storage capacity, which can be added seamlessly on top of..." property="twitter:description">
<meta content="https://linux-git.oraclecorp.com/assets/twitter_card-570ddb06edf56a2312253c5872489847a0f385112ddbcd71ccfa1570febab5d2.jpg" property="twitter:image">

<meta name="csrf-param" content="authenticity_token" />
<meta name="csrf-token" content="lZpQ0nS6Nq6KTr5iwu97muEAcnnksFot8R_l2Jku0aH6L4-ZiVPrlkI0xuttgu8IVcUrYXlUJsgEawgIJNmd6g" />
<meta name="csp-nonce" />
<meta name="action-cable-url" content="/-/cable" />
<link href="/-/manifest.json" rel="manifest">
<link rel="icon" type="image/png" href="/assets/favicon-72a2cad5025aa931d6ea56c3201d1f18e68a8cd39788c7c80d5b2b82aa5143ef.png" id="favicon" data-original-href="/assets/favicon-72a2cad5025aa931d6ea56c3201d1f18e68a8cd39788c7c80d5b2b82aa5143ef.png" />
<link rel="apple-touch-icon" type="image/x-icon" href="/assets/apple-touch-icon-b049d4bc0dd9626f31db825d61880737befc7835982586d015bded10b4435460.png" />
<link href="/search/opensearch.xml" rel="search" title="Search GitLab" type="application/opensearchdescription+xml">




<meta content="Thanos is a set of components that can be composed into a highly available metric system with unlimited storage capacity, which can be added seamlessly on top of..." name="description">
<meta content="#222261" name="theme-color">
</head>

<body class="tab-width-8 gl-browser-generic gl-platform-other" data-group="Prometheus" data-group-full-path="OLCNE/Prometheus" data-namespace-id="1551" data-page="projects:blob:show" data-page-type-id="1e27f1f9b9852379064db1ed41acae54a0b7c495/pkg/basicauth/options.go" data-project="thanos" data-project-full-path="OLCNE/Prometheus/thanos" data-project-id="31330">

<script>
//<![CDATA[
gl = window.gl || {};
gl.client = {"isGeneric":true,"isOther":true};


//]]>
</script>


<header class="header-logged-out" data-testid="navbar">
<a class="gl-sr-only gl-accessibility" href="#content-body">Skip to content</a>
<div class="container-fluid">
<nav aria-label="Explore GitLab" class="header-logged-out-nav gl-flex gl-gap-3 gl-justify-between">
<div class="gl-flex gl-items-center gl-gap-1">
<span class="gl-sr-only">GitLab</span>
<a title="Homepage" id="logo" class="header-logged-out-logo has-tooltip" aria-label="Homepage" href="/"><svg aria-hidden="true" role="img" class="tanuki-logo" width="25" height="24" viewBox="0 0 25 24" fill="none" xmlns="http://www.w3.org/2000/svg">
  <path class="tanuki-shape tanuki" d="m24.507 9.5-.034-.09L21.082.562a.896.896 0 0 0-1.694.091l-2.29 7.01H7.825L5.535.653a.898.898 0 0 0-1.694-.09L.451 9.411.416 9.5a6.297 6.297 0 0 0 2.09 7.278l.012.01.03.022 5.16 3.867 2.56 1.935 1.554 1.176a1.051 1.051 0 0 0 1.268 0l1.555-1.176 2.56-1.935 5.197-3.89.014-.01A6.297 6.297 0 0 0 24.507 9.5Z"
        fill="#E24329"/>
  <path class="tanuki-shape right-cheek" d="m24.507 9.5-.034-.09a11.44 11.44 0 0 0-4.56 2.051l-7.447 5.632 4.742 3.584 5.197-3.89.014-.01A6.297 6.297 0 0 0 24.507 9.5Z"
        fill="#FC6D26"/>
  <path class="tanuki-shape chin" d="m7.707 20.677 2.56 1.935 1.555 1.176a1.051 1.051 0 0 0 1.268 0l1.555-1.176 2.56-1.935-4.743-3.584-4.755 3.584Z"
        fill="#FCA326"/>
  <path class="tanuki-shape left-cheek" d="M5.01 11.461a11.43 11.43 0 0 0-4.56-2.05L.416 9.5a6.297 6.297 0 0 0 2.09 7.278l.012.01.03.022 5.16 3.867 4.745-3.584-7.444-5.632Z"
        fill="#FC6D26"/>
</svg>

</a></div>
<ul class="gl-list-none gl-p-0 gl-m-0 gl-flex gl-gap-3 gl-items-center gl-grow">
<li class="header-logged-out-nav-item">
<a class="" href="/explore">Explore</a>
</li>
</ul>
<ul class="gl-list-none gl-p-0 gl-m-0 gl-flex gl-gap-3 gl-items-center gl-justify-end">
<li class="header-logged-out-nav-item">
<a href="/users/sign_in?redirect_to_referer=yes">Sign in</a>
</li>
</ul>
</nav>
</div>
</header>

<div class="layout-page page-with-super-sidebar">
<aside class="js-super-sidebar super-sidebar super-sidebar-loading" data-command-palette="{&quot;project_files_url&quot;:&quot;/OLCNE/Prometheus/thanos/-/files/1e27f1f9b9852379064db1ed41acae54a0b7c495?format=json&quot;,&quot;project_blob_url&quot;:&quot;/OLCNE/Prometheus/thanos/-/blob/1e27f1f9b9852379064db1ed41acae54a0b7c495&quot;}" data-force-desktop-expanded-sidebar="" data-is-saas="false" data-root-path="/" data-sidebar="{&quot;is_logged_in&quot;:false,&quot;context_switcher_links&quot;:[{&quot;title&quot;:&quot;Explore&quot;,&quot;link&quot;:&quot;/explore&quot;,&quot;icon&quot;:&quot;compass&quot;}],&quot;current_menu_items&quot;:[{&quot;id&quot;:&quot;project_overview&quot;,&quot;title&quot;:&quot;thanos&quot;,&quot;entity_id&quot;:31330,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos&quot;,&quot;link_classes&quot;:&quot;shortcuts-project&quot;,&quot;is_active&quot;:false},{&quot;id&quot;:&quot;manage_menu&quot;,&quot;title&quot;:&quot;Manage&quot;,&quot;icon&quot;:&quot;users&quot;,&quot;avatar_shape&quot;:&quot;rect&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/activity&quot;,&quot;is_active&quot;:false,&quot;items&quot;:[{&quot;id&quot;:&quot;activity&quot;,&quot;title&quot;:&quot;Activity&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/activity&quot;,&quot;link_classes&quot;:&quot;shortcuts-project-activity&quot;,&quot;is_active&quot;:false},{&quot;id&quot;:&quot;members&quot;,&quot;title&quot;:&quot;Members&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/project_members&quot;,&quot;is_active&quot;:false},{&quot;id&quot;:&quot;labels&quot;,&quot;title&quot;:&quot;Labels&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/labels&quot;,&quot;is_active&quot;:false}],&quot;separated&quot;:false},{&quot;id&quot;:&quot;plan_menu&quot;,&quot;title&quot;:&quot;Plan&quot;,&quot;icon&quot;:&quot;planning&quot;,&quot;avatar_shape&quot;:&quot;rect&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/issues&quot;,&quot;is_active&quot;:false,&quot;items&quot;:[{&quot;id&quot;:&quot;project_issue_list&quot;,&quot;title&quot;:&quot;Issues&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/issues&quot;,&quot;pill_count_field&quot;:&quot;openIssuesCount&quot;,&quot;link_classes&quot;:&quot;shortcuts-issues has-sub-items&quot;,&quot;is_active&quot;:false},{&quot;id&quot;:&quot;boards&quot;,&quot;title&quot;:&quot;Issue boards&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/boards&quot;,&quot;link_classes&quot;:&quot;shortcuts-issue-boards&quot;,&quot;is_active&quot;:false},{&quot;id&quot;:&quot;milestones&quot;,&quot;title&quot;:&quot;Milestones&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/milestones&quot;,&quot;is_active&quot;:false},{&quot;id&quot;:&quot;project_wiki&quot;,&quot;title&quot;:&quot;Wiki&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/wikis/home&quot;,&quot;link_classes&quot;:&quot;shortcuts-wiki&quot;,&quot;is_active&quot;:false}],&quot;separated&quot;:false},{&quot;id&quot;:&quot;code_menu&quot;,&quot;title&quot;:&quot;Code&quot;,&quot;icon&quot;:&quot;code&quot;,&quot;avatar_shape&quot;:&quot;rect&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/merge_requests&quot;,&quot;is_active&quot;:true,&quot;items&quot;:[{&quot;id&quot;:&quot;project_merge_request_list&quot;,&quot;title&quot;:&quot;Merge requests&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/merge_requests&quot;,&quot;pill_count_field&quot;:&quot;openMergeRequestsCount&quot;,&quot;link_classes&quot;:&quot;shortcuts-merge_requests&quot;,&quot;is_active&quot;:false},{&quot;id&quot;:&quot;files&quot;,&quot;title&quot;:&quot;Repository&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/tree/1e27f1f9b9852379064db1ed41acae54a0b7c495&quot;,&quot;link_classes&quot;:&quot;shortcuts-tree&quot;,&quot;is_active&quot;:true},{&quot;id&quot;:&quot;branches&quot;,&quot;title&quot;:&quot;Branches&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/branches&quot;,&quot;is_active&quot;:false},{&quot;id&quot;:&quot;commits&quot;,&quot;title&quot;:&quot;Commits&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/commits/1e27f1f9b9852379064db1ed41acae54a0b7c495&quot;,&quot;link_classes&quot;:&quot;shortcuts-commits&quot;,&quot;is_active&quot;:false},{&quot;id&quot;:&quot;tags&quot;,&quot;title&quot;:&quot;Tags&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/tags&quot;,&quot;is_active&quot;:false},{&quot;id&quot;:&quot;graphs&quot;,&quot;title&quot;:&quot;Repository graph&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/network/1e27f1f9b9852379064db1ed41acae54a0b7c495&quot;,&quot;link_classes&quot;:&quot;shortcuts-network&quot;,&quot;is_active&quot;:false},{&quot;id&quot;:&quot;compare&quot;,&quot;title&quot;:&quot;Compare revisions&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/compare?from=oracle-template\u0026to=1e27f1f9b9852379064db1ed41acae54a0b7c495&quot;,&quot;is_active&quot;:false},{&quot;id&quot;:&quot;project_snippets&quot;,&quot;title&quot;:&quot;Snippets&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/snippets&quot;,&quot;link_classes&quot;:&quot;shortcuts-snippets&quot;,&quot;is_active&quot;:false}],&quot;separated&quot;:false},{&quot;id&quot;:&quot;build_menu&quot;,&quot;title&quot;:&quot;Build&quot;,&quot;icon&quot;:&quot;rocket&quot;,&quot;avatar_shape&quot;:&quot;rect&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/pipelines&quot;,&quot;is_active&quot;:false,&quot;items&quot;:[{&quot;id&quot;:&quot;pipelines&quot;,&quot;title&quot;:&quot;Pipelines&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/pipelines&quot;,&quot;link_classes&quot;:&quot;shortcuts-pipelines&quot;,&quot;is_active&quot;:false},{&quot;id&quot;:&quot;jobs&quot;,&quot;title&quot;:&quot;Jobs&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/jobs&quot;,&quot;link_classes&quot;:&quot;shortcuts-builds&quot;,&quot;is_active&quot;:false},{&quot;id&quot;:&quot;pipeline_schedules&quot;,&quot;title&quot;:&quot;Pipeline schedules&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/pipeline_schedules&quot;,&quot;link_classes&quot;:&quot;shortcuts-builds&quot;,&quot;is_active&quot;:false},{&quot;id&quot;:&quot;artifacts&quot;,&quot;title&quot;:&quot;Artifacts&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/artifacts&quot;,&quot;link_classes&quot;:&quot;shortcuts-builds&quot;,&quot;is_active&quot;:false}],&quot;separated&quot;:false},{&quot;id&quot;:&quot;deploy_menu&quot;,&quot;title&quot;:&quot;Deploy&quot;,&quot;icon&quot;:&quot;deployments&quot;,&quot;avatar_shape&quot;:&quot;rect&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/releases&quot;,&quot;is_active&quot;:false,&quot;items&quot;:[{&quot;id&quot;:&quot;releases&quot;,&quot;title&quot;:&quot;Releases&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/releases&quot;,&quot;link_classes&quot;:&quot;shortcuts-deployments-releases&quot;,&quot;is_active&quot;:false},{&quot;id&quot;:&quot;packages_registry&quot;,&quot;title&quot;:&quot;Package registry&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/packages&quot;,&quot;link_classes&quot;:&quot;shortcuts-container-registry&quot;,&quot;is_active&quot;:false},{&quot;id&quot;:&quot;model_registry&quot;,&quot;title&quot;:&quot;Model registry&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/ml/models&quot;,&quot;is_active&quot;:false}],&quot;separated&quot;:false},{&quot;id&quot;:&quot;operations_menu&quot;,&quot;title&quot;:&quot;Operate&quot;,&quot;icon&quot;:&quot;cloud-pod&quot;,&quot;avatar_shape&quot;:&quot;rect&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/environments&quot;,&quot;is_active&quot;:false,&quot;items&quot;:[{&quot;id&quot;:&quot;environments&quot;,&quot;title&quot;:&quot;Environments&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/environments&quot;,&quot;link_classes&quot;:&quot;shortcuts-environments&quot;,&quot;is_active&quot;:false},{&quot;id&quot;:&quot;infrastructure_registry&quot;,&quot;title&quot;:&quot;Terraform modules&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/terraform_module_registry&quot;,&quot;is_active&quot;:false}],&quot;separated&quot;:false},{&quot;id&quot;:&quot;monitor_menu&quot;,&quot;title&quot;:&quot;Monitor&quot;,&quot;icon&quot;:&quot;monitor&quot;,&quot;avatar_shape&quot;:&quot;rect&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/incidents&quot;,&quot;is_active&quot;:false,&quot;items&quot;:[{&quot;id&quot;:&quot;incidents&quot;,&quot;title&quot;:&quot;Incidents&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/incidents&quot;,&quot;is_active&quot;:false}],&quot;separated&quot;:false},{&quot;id&quot;:&quot;analyze_menu&quot;,&quot;title&quot;:&quot;Analyze&quot;,&quot;icon&quot;:&quot;chart&quot;,&quot;avatar_shape&quot;:&quot;rect&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/value_stream_analytics&quot;,&quot;is_active&quot;:false,&quot;items&quot;:[{&quot;id&quot;:&quot;cycle_analytics&quot;,&quot;title&quot;:&quot;Value stream analytics&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/value_stream_analytics&quot;,&quot;link_classes&quot;:&quot;shortcuts-project-cycle-analytics&quot;,&quot;is_active&quot;:false},{&quot;id&quot;:&quot;contributors&quot;,&quot;title&quot;:&quot;Contributor analytics&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/graphs/1e27f1f9b9852379064db1ed41acae54a0b7c495&quot;,&quot;is_active&quot;:false},{&quot;id&quot;:&quot;ci_cd_analytics&quot;,&quot;title&quot;:&quot;CI/CD analytics&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/pipelines/charts&quot;,&quot;is_active&quot;:false},{&quot;id&quot;:&quot;repository_analytics&quot;,&quot;title&quot;:&quot;Repository analytics&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/graphs/1e27f1f9b9852379064db1ed41acae54a0b7c495/charts&quot;,&quot;link_classes&quot;:&quot;shortcuts-repository-charts&quot;,&quot;is_active&quot;:false},{&quot;id&quot;:&quot;model_experiments&quot;,&quot;title&quot;:&quot;Model experiments&quot;,&quot;link&quot;:&quot;/OLCNE/Prometheus/thanos/-/ml/experiments&quot;,&quot;is_active&quot;:false}],&quot;separated&quot;:false}],&quot;current_context_header&quot;:&quot;Project&quot;,&quot;support_path&quot;:&quot;https://about.gitlab.com/get-help/&quot;,&quot;docs_path&quot;:&quot;/help/docs&quot;,&quot;display_whats_new&quot;:false,&quot;show_version_check&quot;:false,&quot;search&quot;:{&quot;search_path&quot;:&quot;/search&quot;,&quot;issues_path&quot;:&quot;/dashboard/issues&quot;,&quot;mr_path&quot;:&quot;/dashboard/merge_requests&quot;,&quot;autocomplete_path&quot;:&quot;/search/autocomplete&quot;,&quot;settings_path&quot;:&quot;/search/settings&quot;,&quot;search_context&quot;:{&quot;group&quot;:{&quot;id&quot;:1551,&quot;name&quot;:&quot;prometheus&quot;,&quot;full_name&quot;:&quot;OLCNE / prometheus&quot;},&quot;group_metadata&quot;:{&quot;issues_path&quot;:&quot;/groups/OLCNE/Prometheus/-/issues&quot;,&quot;mr_path&quot;:&quot;/groups/OLCNE/Prometheus/-/merge_requests&quot;},&quot;project&quot;:{&quot;id&quot;:31330,&quot;name&quot;:&quot;thanos&quot;},&quot;project_metadata&quot;:{&quot;mr_path&quot;:&quot;/OLCNE/Prometheus/thanos/-/merge_requests&quot;,&quot;issues_path&quot;:&quot;/OLCNE/Prometheus/thanos/-/issues&quot;},&quot;code_search&quot;:true,&quot;ref&quot;:&quot;1e27f1f9b9852379064db1ed41acae54a0b7c495&quot;,&quot;scope&quot;:null,&quot;for_snippets&quot;:null}},&quot;panel_type&quot;:&quot;project&quot;,&quot;shortcut_links&quot;:[{&quot;title&quot;:&quot;Snippets&quot;,&quot;href&quot;:&quot;/explore/snippets&quot;,&quot;css_class&quot;:&quot;dashboard-shortcuts-snippets&quot;},{&quot;title&quot;:&quot;Groups&quot;,&quot;href&quot;:&quot;/explore/groups&quot;,&quot;css_class&quot;:&quot;dashboard-shortcuts-groups&quot;},{&quot;title&quot;:&quot;Projects&quot;,&quot;href&quot;:&quot;/explore/projects/starred&quot;,&quot;css_class&quot;:&quot;dashboard-shortcuts-projects&quot;}],&quot;terms&quot;:null}"></aside>

<div class="content-wrapper">
<div class="broadcast-wrapper">




</div>
<div class="alert-wrapper alert-wrapper-top-space gl-flex gl-flex-col gl-gap-3 container-fluid container-limited">
























</div>
<div class="top-bar-fixed container-fluid" data-testid="top-bar">
<div class="top-bar-container gl-flex gl-items-center gl-gap-2">
<div class="gl-grow gl-basis-0 gl-flex gl-items-center gl-justify-start gl-gap-3">
<button class="gl-button btn btn-icon btn-md btn-default btn-default-tertiary js-super-sidebar-toggle-expand super-sidebar-toggle -gl-ml-3" aria-controls="super-sidebar" aria-expanded="false" aria-label="Primary navigation sidebar" type="button"><svg class="s16 gl-icon gl-button-icon " data-testid="sidebar-icon"><use href="/assets/icons-aa2c8ddf99d22b77153ca2bb092a23889c12c597fc8b8de94b0f730eb53513f6.svg#sidebar"></use></svg>

</button>
<script type="application/ld+json">
{"@context":"https://schema.org","@type":"BreadcrumbList","itemListElement":[{"@type":"ListItem","position":1,"name":"OLCNE","item":"https://linux-git.oraclecorp.com/OLCNE"},{"@type":"ListItem","position":2,"name":"prometheus","item":"https://linux-git.oraclecorp.com/OLCNE/Prometheus"},{"@type":"ListItem","position":3,"name":"thanos","item":"https://linux-git.oraclecorp.com/OLCNE/Prometheus/thanos"},{"@type":"ListItem","position":4,"name":"Repository","item":"https://linux-git.oraclecorp.com/OLCNE/Prometheus/thanos/-/blob/1e27f1f9b9852379064db1ed41acae54a0b7c495/pkg/basicauth/options.go"}]}


</script>
<div data-testid="breadcrumb-links" id="js-vue-page-breadcrumbs-wrapper">
<div data-breadcrumbs-json="[{&quot;text&quot;:&quot;OLCNE&quot;,&quot;href&quot;:&quot;/OLCNE&quot;,&quot;avatarPath&quot;:&quot;/uploads/-/system/group/avatar/550/ocne.png&quot;},{&quot;text&quot;:&quot;prometheus&quot;,&quot;href&quot;:&quot;/OLCNE/Prometheus&quot;,&quot;avatarPath&quot;:null},{&quot;text&quot;:&quot;thanos&quot;,&quot;href&quot;:&quot;/OLCNE/Prometheus/thanos&quot;,&quot;avatarPath&quot;:null},{&quot;text&quot;:&quot;Repository&quot;,&quot;href&quot;:&quot;/OLCNE/Prometheus/thanos/-/blob/1e27f1f9b9852379064db1ed41acae54a0b7c495/pkg/basicauth/options.go&quot;,&quot;avatarPath&quot;:null}]" id="js-vue-page-breadcrumbs"></div>
<div id="js-injected-page-breadcrumbs"></div>
</div>


</div>
<div class="gl-flex-none gl-flex gl-items-center gl-justify-center">
<div id="js-advanced-search-modal"></div>

</div>
<div class="gl-grow gl-basis-0 gl-flex gl-items-center gl-justify-end">
<div id="js-work-item-feedback"></div>


</div>
</div>
</div>

<div class="container-fluid container-limited project-highlight-puc">
<main class="content" id="content-body" itemscope itemtype="http://schema.org/SoftwareSourceCode">
<div class="flash-container flash-container-page sticky" data-testid="flash-container">
<div id="js-global-alerts"></div>
</div>






<div class="js-signature-container" data-signatures-path="/OLCNE/Prometheus/thanos/-/commits/60e5ae4d2b5448693bc8ce6035919900f7ac8431/signatures?limit=1"></div>

<div class="tree-holder gl-pt-4" id="tree-holder">
<div class="nav-block">
<div class="tree-ref-container">
<div class="tree-ref-holder gl-max-w-26">
<div data-project-id="31330" data-project-root-path="/OLCNE/Prometheus/thanos" data-ref="1e27f1f9b9852379064db1ed41acae54a0b7c495" data-ref-type="" id="js-tree-ref-switcher"></div>
</div>
<ul class="breadcrumb repo-breadcrumb">
<li class="breadcrumb-item">
<a href="/OLCNE/Prometheus/thanos/-/tree/1e27f1f9b9852379064db1ed41acae54a0b7c495">thanos
</a></li>
<li class="breadcrumb-item">
<a href="/OLCNE/Prometheus/thanos/-/tree/1e27f1f9b9852379064db1ed41acae54a0b7c495/pkg">pkg</a>
</li>
<li class="breadcrumb-item">
<a href="/OLCNE/Prometheus/thanos/-/tree/1e27f1f9b9852379064db1ed41acae54a0b7c495/pkg/basicauth">basicauth</a>
</li>
<li class="breadcrumb-item">
<a href="/OLCNE/Prometheus/thanos/-/blob/1e27f1f9b9852379064db1ed41acae54a0b7c495/pkg/basicauth/options.go"><strong>options.go</strong>
</a></li>
</ul>
</div>
<div class="tree-controls gl-flex gl-flex-wrap sm:gl-flex-nowrap gl-items-baseline gl-gap-3">
<button class="gl-button btn btn-md btn-default has-tooltip shortcuts-find-file" title="Go to file, press &lt;kbd class=&#39;flat ml-1&#39; aria-hidden=true&gt;/~&lt;/kbd&gt; or &lt;kbd class=&#39;flat ml-1&#39; aria-hidden=true&gt;t&lt;/kbd&gt;" aria-keyshortcuts="/+~ t" data-html="true" data-event-tracking="click_find_file_button_on_repository_pages" type="button"><span class="gl-button-text">
Find file

</span>

</button>
<a data-event-tracking="click_blame_control_on_blob_page" class="gl-button btn btn-md btn-default js-blob-blame-link" href="/OLCNE/Prometheus/thanos/-/blame/1e27f1f9b9852379064db1ed41acae54a0b7c495/pkg/basicauth/options.go"><span class="gl-button-text">
Blame
</span>

</a>
<a aria-keyshortcuts="y" class="gl-button btn btn-md btn-default has-tooltip js-data-file-blob-permalink-url" data-html="true" title="Go to permalink &lt;kbd class=&#39;flat ml-1&#39; aria-hidden=true&gt;y&lt;/kbd&gt;" href="/OLCNE/Prometheus/thanos/-/blob/1e27f1f9b9852379064db1ed41acae54a0b7c495/pkg/basicauth/options.go"><span class="gl-button-text">
Permalink
</span>

</a>
</div>
</div>

<div class="info-well">
<div class="well-segment">
<ul class="blob-commit-info">
<li class="commit flex-row js-toggle-container" id="commit-60e5ae4d">
<div class="gl-self-start gl-block">
<a href="/fred.tibbitts"><img alt="Fred Tibbitts III&#39;s avatar" src="/assets/no_avatar-849f9c04a3a0d0cea2424ae97b27447dc64a7dbfae83c036c45b403392f0e8ba.png" class="avatar s32 gl-inline-block" title="Fred Tibbitts III"></a>
</div>
<div class="commit-detail flex-list gl-flex gl-justify-between gl-items-start gl-grow gl-min-w-0">
<div class="commit-content gl-self-center" data-testid="commit-content">
<div class="gl-flex sm:gl-hidden gl-gap-3 gl-items-center">
<div class="committer gl-text-sm">
<time class="js-timeago" title="Sep 11, 2023 12:03pm" datetime="2023-09-11T19:03:24Z" tabindex="0" aria-label="Sep 11, 2023 12:03pm" data-toggle="tooltip" data-placement="bottom" data-container="body">Sep 11, 2023</time>
</div>
<a class="gl-button btn btn-md btn-link commit-row-message js-onboarding-commit-item" href="/OLCNE/Prometheus/thanos/-/commit/60e5ae4d2b5448693bc8ce6035919900f7ac8431"><svg class="s16 gl-icon gl-button-icon " data-testid="commit-icon"><use href="/assets/icons-aa2c8ddf99d22b77153ca2bb092a23889c12c597fc8b8de94b0f730eb53513f6.svg#commit"></use></svg>
<span class="gl-button-text">
60e5ae4d

</span>

</a></div>
<div class="gl-hidden sm:gl-block">
<a class="commit-row-message item-title js-onboarding-commit-item " href="/OLCNE/Prometheus/thanos/-/commit/60e5ae4d2b5448693bc8ce6035919900f7ac8431">Changes to create release 0.32.2</a>
<span class="commit-row-message d-inline d-sm-none">
&middot;
60e5ae4d
</span>
<div class="committer gl-text-sm">
<a class="commit-author-link js-user-link" data-user-id="1585" href="/fred.tibbitts">Fred Tibbitts III</a> authored <time class="js-timeago" title="Sep 11, 2023 12:03pm" datetime="2023-09-11T19:03:24Z" tabindex="0" aria-label="Sep 11, 2023 12:03pm" data-toggle="tooltip" data-placement="bottom" data-container="body">Sep 11, 2023</time>
</div>


</div>
</div>
<div class="commit-actions gl-flex gl-items-center gl-gap-3">
<div class="gl-hidden sm:gl-flex gl-items-center gl-gap-3">

<div class="js-commit-pipeline-status" data-endpoint="/OLCNE/Prometheus/thanos/-/commit/60e5ae4d2b5448693bc8ce6035919900f7ac8431/pipelines?ref=1e27f1f9b9852379064db1ed41acae54a0b7c495"></div>
<div class="btn-group gl-hidden sm:gl-flex">
<span class="gl-button btn btn-label btn-md btn-default dark:!gl-bg-neutral-800" type="button"><span class="gl-button-text gl-font-monospace">
60e5ae4d

</span>

</span><button class="gl-button btn btn-icon btn-md btn-default " title="Copy commit SHA" aria-label="Copy commit SHA" aria-live="polite" data-toggle="tooltip" data-placement="bottom" data-container="body" data-html="true" data-category="primary" data-size="medium" data-clipboard-text="60e5ae4d2b5448693bc8ce6035919900f7ac8431" type="button"><svg class="s16 gl-icon gl-button-icon " data-testid="copy-to-clipboard-icon"><use href="/assets/icons-aa2c8ddf99d22b77153ca2bb092a23889c12c597fc8b8de94b0f730eb53513f6.svg#copy-to-clipboard"></use></svg>

</button>

</div>
</div>
<div class="gl-block sm:gl-hidden">
<button class="gl-button btn btn-icon btn-md btn-default button-ellipsis-horizontal text-expander js-toggle-button" data-toggle="tooltip" data-container="body" data-collapse-title="Toggle commit description" data-expand-title="Toggle commit description" title="Toggle commit description" aria-label="Toggle commit description" type="button"><svg class="s16 gl-icon gl-button-icon " data-testid="ellipsis_h-icon"><use href="/assets/icons-aa2c8ddf99d22b77153ca2bb092a23889c12c597fc8b8de94b0f730eb53513f6.svg#ellipsis_h"></use></svg>

</button>
</div>
<div data-event-tracking="click_history_control_on_blob_page" data-history-link="/OLCNE/Prometheus/thanos/-/commits/1e27f1f9b9852379064db1ed41acae54a0b7c495/pkg/basicauth/options.go" id="js-commit-history-link"></div>
</div>
</div>
<div class="gl-block sm:gl-hidden">
<div class="gl-hidden js-toggle-content gl-mt-6">
<a class="commit-row-message item-title js-onboarding-commit-item " href="/OLCNE/Prometheus/thanos/-/commit/60e5ae4d2b5448693bc8ce6035919900f7ac8431">Changes to create release 0.32.2</a>
<div class="committer gl-text-sm">
<a class="commit-author-link js-user-link" data-user-id="1585" href="/fred.tibbitts">Fred Tibbitts III</a> authored <time class="js-timeago" title="Sep 11, 2023 12:03pm" datetime="2023-09-11T19:03:24Z" tabindex="0" aria-label="Sep 11, 2023 12:03pm" data-toggle="tooltip" data-placement="bottom" data-container="body">Sep 11, 2023</time>
</div>

</div>
</div>
</li>

</ul>
</div>
<div class="gl-hidden sm:gl-block">

</div>
</div>
<div class="blob-content-holder js-per-page" data-blame-per-page="1000" id="blob-content-holder">
<div data-blob-path="pkg/basicauth/options.go" data-can-download-code="true" data-original-branch="1e27f1f9b9852379064db1ed41acae54a0b7c495" data-project-path="OLCNE/Prometheus/thanos" data-ref-type="" data-resource-id="gid://gitlab/Project/31330" data-user-id="" id="js-view-blob-app">
<div class="gl-spinner-container" role="status"><span aria-hidden class="gl-spinner gl-spinner-md gl-spinner-dark !gl-align-text-bottom"></span><span class="gl-sr-only !gl-absolute">Loading</span>
</div>
</div>
</div>

</div>
<script>
//<![CDATA[
  window.gl = window.gl || {};
  window.gl.webIDEPath = '/-/ide/project/OLCNE/Prometheus/thanos/edit/1e27f1f9b9852379064db1ed41acae54a0b7c495/-/pkg/basicauth/options.go'


//]]>
</script>
<div data-ambiguous="false" data-ref="1e27f1f9b9852379064db1ed41acae54a0b7c495" id="js-ambiguous-ref-modal"></div>

</main>
</div>


</div>
</div>


<script>
//<![CDATA[
if ('loading' in HTMLImageElement.prototype) {
  document.querySelectorAll('img.lazy').forEach(img => {
    img.loading = 'lazy';
    let imgUrl = img.dataset.src;
    // Only adding width + height for avatars for now
    if (imgUrl.indexOf('/avatar/') > -1 && imgUrl.indexOf('?') === -1) {
      const targetWidth = img.getAttribute('width') || img.width;
      imgUrl += `?width=${targetWidth}`;
    }
    img.src = imgUrl;
    img.removeAttribute('data-src');
    img.classList.remove('lazy');
    img.classList.add('js-lazy-loaded');
    img.dataset.testid = 'js-lazy-loaded-content';
  });
}

//]]>
</script>
<script>
//<![CDATA[
gl = window.gl || {};
gl.experiments = {};


//]]>
</script>

</body>
</html>

