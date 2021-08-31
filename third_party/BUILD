go_module(
    name = "github.com.google.go-github",
    module = "github.com/google/go-github",
    version = "v17.0.0+incompatible",
    install = [
        "github",
    ],
    deps = [
        ":github.com.google.go-querystring",
    ],
)

go_module(
    name = "github.com.aws.aws-sdk-go-v2.service.internal.s3shared",
    module = "github.com/aws/aws-sdk-go-v2/service/internal/s3shared",
    version = "v1.2.2",
    install = [
        ".",
        "config",
        "arn",
    ],
    deps = [
        ":github.com.aws.smithy-go",
        ":github.com.aws.aws-sdk-go-v2",
    ],
)

go_module(
    name = "github.com.containerd.continuity",
    module = "github.com/containerd/continuity",
    version = "v0.0.0-20200710164510-efbc4488d8fe",
    install = [
        "pathdriver",
        "fs",
        "sysx",
    ],
    deps = [
        ":github.com.pkg.errors",
        ":github.com.sirupsen.logrus",
        ":golang.org.x.sys",
        ":golang.org.x.sync",
    ],
)

go_module(
    name = "github.com.opencontainers.runc",
    module = "github.com/opencontainers/runc",
    version = "v1.0.0-rc9",
    install = [
        "libcontainer/user",
        "libcontainer/system",
    ],
    deps = [
        ":golang.org.x.sys",
    ],
)

go_module(
    name = "github.com.Azure.go-autorest.autorest.validation",
    module = "github.com/Azure/go-autorest/autorest/validation",
    version = "v0.3.1",
)

go_module(
    name = "github.com.coreos.go-systemd",
    module = "github.com/coreos/go-systemd",
    version = "v0.0.0-20191104093116-d3cd4ed1dbcf",
    install = [
        "daemon",
    ],
)

go_module(
    name = "github.com.hashicorp.vault-plugin-auth-cf",
    module = "github.com/hashicorp/vault-plugin-auth-cf",
    version = "v0.9.0",
    install = [
        ".",
        "signatures",
        "models",
        "util",
    ],
    deps = [
        ":github.com.cloudfoundry-community.go-cfclient",
        ":github.com.hashicorp.vault.sdk_1",
        ":github.com.pkg.errors",
        ":github.com.hashicorp.vault.api",
        ":github.com.hashicorp.go-multierror",
        ":github.com.hashicorp.go-sockaddr",
        ":github.com.hashicorp.go-cleanhttp",
    ],
)

go_module(
    name = "github.com.Azure.azure-sdk-for-go",
    module = "github.com/Azure/azure-sdk-for-go",
    version = "v51.1.0+incompatible",
    install = [
        "services/keyvault/v7.0/keyvault",
        "services/preview/authorization/mgmt/2018-01-01-preview/authorization",
        "services/network/mgmt/2015-06-15/network",
        "services/graphrbac/1.6/graphrbac",
        "version",
        "services/compute/mgmt/2019-07-01/compute",
    ],
    deps = [
        ":github.com.Azure.go-autorest.tracing",
        ":github.com.Azure.go-autorest.autorest",
        ":github.com.Azure.go-autorest.autorest.to",
        ":github.com.Azure.go-autorest.autorest.validation",
        ":github.com.Azure.go-autorest.autorest.date",
    ],
)

go_module(
    name = "github.com.hashicorp.raft-snapshot",
    module = "github.com/hashicorp/raft-snapshot",
    version = "v1.0.3",
    deps = [
        ":github.com.hashicorp.go-hclog",
        ":github.com.hashicorp.raft",
    ],
)

go_module(
    name = "github.com.coreos.go-oidc",
    module = "github.com/coreos/go-oidc",
    version = "v2.2.1+incompatible",
    deps = [
        ":golang.org.x.oauth2",
        ":github.com.pquerna.cachecontrol",
        ":gopkg.in.square.go-jose.v2",
    ],
)

go_module(
    name = "github.com.kr.pretty",
    module = "github.com/kr/pretty",
    version = "v0.2.1",
    deps = [
        ":github.com.kr.text",
    ],
)

go_module(
    name = "github.com.docker.docker",
    module = "github.com/docker/docker",
    version = "v17.12.0-ce-rc1.0.20200309214505-aa6a9891b09c+incompatible",
    install = [
        "pkg/pools",
        "errdefs",
        "api/types/mount",
        "pkg/fileutils",
        "client",
        "api/types/blkiodev",
        "pkg/idtools",
        "api/types/network",
        "pkg/archive",
        "api/types/registry",
        "api/types/swarm",
        "pkg/system",
        "api/types/events",
        "pkg/mount",
        "api/types/filters",
        "api/types/time",
        "api/types/strslice",
        "api/types/container",
        "api/types/image",
        "api/types/swarm/runtime",
        "pkg/ioutils",
        "api/types/volume",
        "api/types/versions",
        "api",
        "api/types",
    ],
    deps = [
        ":github.com.containerd.containerd",
        ":github.com.docker.distribution",
        ":github.com.sirupsen.logrus",
        ":google.golang.org.grpc",
        ":github.com.opencontainers.go-digest",
        ":github.com.docker.go-connections",
        ":github.com.pkg.errors",
        ":github.com.opencontainers.runc",
        ":github.com.containerd.continuity",
        ":golang.org.x.sys",
        ":github.com.opencontainers.image-spec",
        ":github.com.docker.go-units",
        ":github.com.gogo.protobuf",
    ],
)

go_module(
    name = "github.com.DataDog.datadog-go",
    module = "github.com/DataDog/datadog-go",
    version = "v3.2.0+incompatible",
    install = [
        "statsd",
    ],
)

go_module(
    name = "github.com.mitchellh.cli",
    module = "github.com/mitchellh/cli",
    version = "v1.1.2",
    deps = [
        ":github.com.armon.go-radix",
        ":github.com.Masterminds.sprig",
        ":github.com.posener.complete",
        ":github.com.fatih.color",
        ":github.com.mattn.go-isatty",
        ":github.com.bgentry.speakeasy",
    ],
)

go_module(
    name = "github.com.tklauser.go-sysconf",
    module = "github.com/tklauser/go-sysconf",
    version = "v0.3.6",
    deps = [
        ":golang.org.x.sys",
        ":github.com.tklauser.numcpus",
    ],
)

go_module(
    name = "github.com.renier.xmlrpc",
    module = "github.com/renier/xmlrpc",
    version = "v0.0.0-20170708154548-ce4a1a486c03",
    deps = [
        ":golang.org.x.text",
    ],
)

go_module(
    name = "golang.org.x.time",
    module = "golang.org/x/time",
    version = "v0.0.0-20210723032227-1f47c861a9ac",
    install = [
        "rate",
    ],
)

go_module(
    name = "github.com.hashicorp.vault.api",
    module = "github.com/hashicorp/vault/api",
    version = "v1.1.2-0.20210713235431-1fc8af4c041f",
    deps = [
        ":github.com.hashicorp.go-cleanhttp",
        ":github.com.cenkalti.backoff.v3",
        ":github.com.hashicorp.hcl",
        ":github.com.hashicorp.go-multierror",
        ":github.com.hashicorp.vault.sdk_1",
        ":golang.org.x.time",
        ":github.com.hashicorp.go-rootcerts",
        ":github.com.hashicorp.errwrap",
        ":golang.org.x.net",
        ":gopkg.in.square.go-jose.v2",
        ":github.com.mitchellh.mapstructure",
        ":github.com.hashicorp.go-retryablehttp",
    ],
)

go_module(
    name = "golang.org.x.text",
    module = "golang.org/x/text",
    version = "v0.3.6",
    install = [
        "internal/language/compact",
        "internal/utf8internal",
        "encoding/unicode",
        "unicode/bidi",
        "internal/tag",
        "internal/language",
        "runes",
        "encoding",
        "encoding/internal/identifier",
        "transform",
        "encoding/internal",
        "encoding/charmap",
        "secure/bidirule",
        "language",
        "unicode/norm",
    ],
)

go_module(
    name = "github.com.oracle.oci-go-sdk",
    module = "github.com/oracle/oci-go-sdk",
    version = "v12.5.0+incompatible",
    install = [
        "common/auth",
        "common",
        "keymanagement",
        "objectstorage",
    ],
)

go_module(
    name = "github.com.influxdata.influxdb",
    module = "github.com/influxdata/influxdb",
    version = "v0.0.0-20190411212539-d24b7ba8c4c4",
    install = [
        "pkg/escape",
        "models",
        "client/v2",
    ],
)

go_module(
    name = "github.com.hashicorp.go-tfe",
    module = "github.com/hashicorp/go-tfe",
    version = "v0.12.0",
    deps = [
        ":github.com.google.go-querystring",
        ":golang.org.x.time",
        ":github.com.svanharmelen.jsonapi",
        ":github.com.hashicorp.go-cleanhttp",
        ":github.com.hashicorp.go-retryablehttp",
        ":github.com.hashicorp.go-slug",
    ],
)

go_module(
    name = "go.uber.org.zap",
    module = "go.uber.org/zap",
    version = "v1.14.1",
    install = [
        ".",
        "buffer",
        "zapcore",
        "internal/color",
        "internal/exit",
        "internal/bufferpool",
    ],
    deps = [
        ":go.uber.org.multierr",
        ":go.uber.org.atomic",
    ],
)

go_module(
    name = "github.com.go-logr.logr",
    module = "github.com/go-logr/logr",
    version = "v0.4.0",
)

go_module(
    name = "github.com.Sectorbob.mlab-ns2",
    module = "github.com/Sectorbob/mlab-ns2",
    version = "v0.0.0-20171030222938-d3aa0c295a8a",
    install = [
        "gae/ns/digest",
    ],
)

go_module(
    name = "github.com.aws.smithy-go",
    module = "github.com/aws/smithy-go",
    version = "v1.3.1",
    install = [
        ".",
        "logging",
        "rand",
        "encoding/httpbinding",
        "encoding/xml",
        "encoding",
        "transport/http",
        "middleware",
        "waiter",
        "io",
        "ptr",
        "time",
        "transport/http/internal/io",
    ],
)

go_module(
    name = "github.com.docker.go-units",
    module = "github.com/docker/go-units",
    version = "v0.4.0",
)

go_module(
    name = "github.com.hashicorp.go-retryablehttp",
    module = "github.com/hashicorp/go-retryablehttp",
    version = "v0.6.7",
    deps = [
        ":github.com.hashicorp.go-cleanhttp",
    ],
)

go_module(
    name = "github.com.jcmturner.rpc.v2",
    module = "github.com/jcmturner/rpc/v2",
    version = "v2.0.2",
    install = [
        "mstypes",
        "ndr",
    ],
    deps = [
        ":golang.org.x.net",
    ],
)

go_module(
    name = "github.com.prometheus.client_model",
    module = "github.com/prometheus/client_model",
    version = "v0.2.0",
    install = [
        "go",
    ],
    deps = [
        ":github.com.golang.protobuf",
    ],
)

go_module(
    name = "google.golang.org.protobuf",
    module = "google.golang.org/protobuf",
    version = "v1.26.0",
    install = [
        "types/known/emptypb",
        "internal/flags",
        "internal/order",
        "internal/errors",
        "internal/strs",
        "reflect/protodesc",
        "internal/descfmt",
        "internal/filedesc",
        "proto",
        "internal/descopts",
        "encoding/protojson",
        "internal/impl",
        "internal/encoding/messageset",
        "reflect/protoreflect",
        "types/known/wrapperspb",
        "internal/encoding/text",
        "internal/pragma",
        "internal/filetype",
        "internal/encoding/defval",
        "reflect/protoregistry",
        "encoding/prototext",
        "encoding/protowire",
        "types/known/structpb",
        "types/descriptorpb",
        "internal/encoding/json",
        "runtime/protoimpl",
        "types/known/anypb",
        "types/known/fieldmaskpb",
        "internal/version",
        "internal/set",
        "runtime/protoiface",
        "types/known/timestamppb",
        "types/known/durationpb",
        "internal/encoding/tag",
        "internal/genid",
        "internal/detrand",
    ],
)

go_module(
    name = "github.com.Azure.azure-pipeline-go",
    module = "github.com/Azure/azure-pipeline-go",
    version = "v0.2.3",
    install = [
        "pipeline",
    ],
    deps = [
        ":github.com.mattn.go-ieproxy",
    ],
)

go_module(
    name = "github.com.hashicorp.vault-plugin-secrets-ad",
    module = "github.com/hashicorp/vault-plugin-secrets-ad",
    version = "v0.10.0",
    install = [
        "plugin/client",
        "plugin/util",
        "plugin",
    ],
    deps = [
        ":github.com.go-errors.errors",
        ":github.com.hashicorp.go-hclog",
        ":github.com.hashicorp.vault.sdk_1",
        ":golang.org.x.text",
        ":github.com.go-ldap.ldap.v3",
        ":github.com.patrickmn.go-cache",
        ":github.com.mitchellh.mapstructure",
        ":github.com.armon.go-metrics",
    ],
)

go_module(
    name = "github.com.pquerna.cachecontrol",
    module = "github.com/pquerna/cachecontrol",
    version = "v0.0.0-20201205024021-ac21108117ac",
    install = [
        ".",
        "cacheobject",
    ],
)

go_module(
    name = "github.com.xi2.xz",
    module = "github.com/xi2/xz",
    version = "v0.0.0-20171230120015-48954b6210f8",
)

go_module(
    name = "github.com.mattn.go-shellwords",
    module = "github.com/mattn/go-shellwords",
    version = "v1.0.12",
)

go_module(
    name = "github.com.hashicorp.go-plugin",
    module = "github.com/hashicorp/go-plugin",
    version = "v1.0.1",
    install = [
        ".",
        "internal/plugin",
    ],
    deps = [
        ":github.com.mitchellh.go-testing-interface",
        ":golang.org.x.net",
        ":github.com.hashicorp.go-hclog",
        ":github.com.hashicorp.yamux",
        ":google.golang.org.grpc",
        ":github.com.oklog.run",
        ":github.com.golang.protobuf",
    ],
)

go_module(
    name = "github.com.googleapis.gnostic",
    module = "github.com/googleapis/gnostic",
    version = "v0.5.5",
    install = [
        "openapiv2",
        "extensions",
        "jsonschema",
        "compiler",
    ],
    deps = [
        ":google.golang.org.protobuf",
        ":gopkg.in.yaml.v3",
        ":github.com.golang.protobuf",
    ],
)

go_module(
    name = "github.com.miekg.dns",
    module = "github.com/miekg/dns",
    version = "v1.1.40",
    deps = [
        ":golang.org.x.crypto",
        ":golang.org.x.net",
        ":golang.org.x.sys",
    ],
)

go_module(
    name = "github.com.hashicorp.vic",
    module = "github.com/hashicorp/vic",
    version = "v1.5.1-0.20190403131502-bbfe86ec9443",
    install = [
        "pkg/vsphere/tags",
    ],
    deps = [
        ":github.com.sirupsen.logrus",
        ":github.com.vmware.govmomi",
        ":github.com.pkg.errors",
    ],
)

go_module(
    name = "github.com.gophercloud.gophercloud",
    module = "github.com/gophercloud/gophercloud",
    version = "v0.1.0",
    install = [
        ".",
        "openstack",
        "openstack/compute/v2/images",
        "openstack/identity/v2/tenants",
        "openstack/identity/v2/tokens",
        "openstack/compute/v2/servers",
        "pagination",
        "openstack/compute/v2/flavors",
        "openstack/utils",
        "openstack/identity/v3/tokens",
    ],
)

go_module(
    name = "github.com.couchbase.gocbcore.v9",
    module = "github.com/couchbase/gocbcore/v9",
    version = "v9.0.4",
    install = [
        ".",
        "connstr",
        "memd",
        "scram",
    ],
    deps = [
        ":github.com.google.uuid",
        ":github.com.golang.snappy",
    ],
)

go_module(
    name = "k8s.io.utils",
    module = "k8s.io/utils",
    version = "v0.0.0-20210707171843-4b05e18ac7d9",
    install = [
        "buffer",
        "integer",
        "trace",
        "pointer",
    ],
    deps = [
        ":k8s.io.klog.v2",
    ],
)

go_module(
    name = "github.com.davecgh.go-spew",
    module = "github.com/davecgh/go-spew",
    version = "v1.1.1",
    install = [
        "spew",
    ],
)

go_module(
    name = "gopkg.in.mgo.v2",
    module = "gopkg.in/mgo.v2",
    version = "v2.0.0-20180705113604-9856a29383ce",
    install = [
        ".",
        "internal/json",
        "bson",
        "internal/scram",
    ],
)

go_module(
    name = "github.com.boombuler.barcode",
    module = "github.com/boombuler/barcode",
    version = "v1.0.1-0.20190219062509-6c824513bacc",
    install = [
        ".",
        "utils",
        "qr",
    ],
)

go_module(
    name = "github.com.hashicorp.vault-plugin-secrets-gcpkms",
    module = "github.com/hashicorp/vault-plugin-secrets-gcpkms",
    version = "v0.9.0",
    deps = [
        ":github.com.golang.protobuf",
        ":github.com.jeffchao.backoff",
        ":google.golang.org.genproto",
        ":github.com.hashicorp.vault.sdk_1",
        ":cloud.google.com.go",
        ":github.com.hashicorp.errwrap",
        ":github.com.gammazero.workerpool",
        ":google.golang.org.grpc",
        ":github.com.hashicorp.go-multierror",
        ":golang.org.x.oauth2",
        ":google.golang.org.api",
        ":google.golang.org.genproto_1",
    ],
)

go_module(
    name = "github.com.snowflakedb.gosnowflake",
    module = "github.com/snowflakedb/gosnowflake",
    version = "v1.6.1",
    deps = [
        ":github.com.apache.arrow.go.arrow",
        ":github.com.form3tech-oss.jwt-go",
        ":github.com.aws.aws-sdk-go-v2.service.s3",
        ":github.com.sirupsen.logrus",
        ":golang.org.x.crypto",
        ":github.com.aws.aws-sdk-go-v2",
        ":github.com.aws.smithy-go",
        ":github.com.aws.aws-sdk-go-v2.credentials",
        ":github.com.aws.aws-sdk-go-v2.feature.s3.manager",
        ":github.com.google.uuid",
        ":github.com.Azure.azure-storage-blob-go",
        ":github.com.pkg.browser",
    ],
)

go_module(
    name = "github.com.apache.arrow.go.arrow",
    module = "github.com/apache/arrow/go/arrow",
    version = "v0.0.0-20200601151325-b2287a20f230",
    install = [
        ".",
        "decimal128",
        "array",
        "arrio",
        "internal/debug",
        "internal/cpu",
        "ipc",
        "memory",
        "bitutil",
        "float16",
        "internal/flatbuf",
    ],
    deps = [
        ":golang.org.x.xerrors",
        ":github.com.google.flatbuffers",
    ],
)

go_module(
    name = "github.com.tv42.httpunix",
    module = "github.com/tv42/httpunix",
    version = "v0.0.0-20150427012821-b75d8614f926",
)

go_module(
    name = "github.com.pkg.browser",
    module = "github.com/pkg/browser",
    version = "v0.0.0-20180916011732-0a3d74bf9ce4",
)

go_module(
    name = "github.com.packethost.packngo",
    module = "github.com/packethost/packngo",
    version = "v0.1.1-0.20180711074735-b9cb5096f54c",
)

go_module(
    name = "github.com.hashicorp.consul.api",
    module = "github.com/hashicorp/consul/api",
    version = "v1.4.0",
    deps = [
        ":github.com.hashicorp.serf",
        ":github.com.hashicorp.go-rootcerts",
        ":github.com.hashicorp.go-cleanhttp",
        ":github.com.hashicorp.go-hclog",
        ":github.com.mitchellh.mapstructure",
    ],
)

go_module(
    name = "github.com.hashicorp.vault-plugin-auth-centrify",
    module = "github.com/hashicorp/vault-plugin-auth-centrify",
    version = "v0.9.0",
    deps = [
        ":github.com.hashicorp.vault.api",
        ":github.com.hashicorp.vault.sdk_1",
        ":github.com.centrify.cloud-golang-sdk",
        ":github.com.hashicorp.go-cleanhttp",
    ],
)

go_module(
    name = "github.com.spf13.pflag",
    module = "github.com/spf13/pflag",
    version = "v1.0.5",
)

go_module(
    name = "google.golang.org.api",
    module = "google.golang.org/api",
    version = "v0.29.0",
    install = [
        "cloudresourcemanager/v1",
        "transport/http/internal/propagation",
        "internal/gensupport",
        "compute/v1",
        "transport/http",
        "transport/cert",
        "option/internaloption",
        "option",
        "transport/grpc",
        "storage/v1",
        "internal/third_party/uritemplates",
        "internal",
        "googleapi/transport",
        "admin/directory/v1",
        "googleapi",
        "iterator",
        "iam/v1",
        "iamcredentials/v1",
    ],
    deps = [
        ":go.opencensus.io",
        ":github.com.googleapis.gax-go.v2",
        ":golang.org.x.oauth2",
        ":google.golang.org.grpc",
        ":golang.org.x.sys",
    ],
)

go_module(
    name = "github.com.prometheus.client_golang",
    module = "github.com/prometheus/client_golang",
    version = "v1.7.1",
    install = [
        "prometheus",
        "prometheus/internal",
        "prometheus/push",
    ],
    deps = [
        ":github.com.golang.protobuf",
        ":github.com.beorn7.perks",
        ":github.com.prometheus.common",
        ":github.com.prometheus.procfs",
        ":github.com.prometheus.client_model",
        ":github.com.cespare.xxhash.v2",
    ],
)

go_module(
    name = "github.com.nicolai86.scaleway-sdk",
    module = "github.com/nicolai86/scaleway-sdk",
    version = "v1.10.2-0.20180628010248-798f60e20bb2",
    deps = [
        ":golang.org.x.sync",
    ],
)

go_module(
    name = "github.com.pmezard.go-difflib",
    module = "github.com/pmezard/go-difflib",
    version = "v1.0.0",
    install = [
        "difflib",
    ],
)

go_module(
    name = "gopkg.in.yaml.v3",
    module = "gopkg.in/yaml.v3",
    version = "v3.0.0-20210107192922-496545a6307b",
)

go_module(
    name = "github.com.hashicorp.vault-plugin-auth-alicloud",
    module = "github.com/hashicorp/vault-plugin-auth-alicloud",
    version = "v0.9.0",
    install = [
        ".",
        "tools",
    ],
    deps = [
        ":github.com.aliyun.alibaba-cloud-sdk-go",
        ":github.com.hashicorp.vault.api",
        ":github.com.hashicorp.errwrap",
        ":github.com.hashicorp.vault.sdk_1",
        ":github.com.hashicorp.go-cleanhttp",
        ":github.com.hashicorp.go-sockaddr",
    ],
)

go_module(
    name = "go.opencensus.io",
    module = "go.opencensus.io",
    version = "v0.22.3",
    install = [
        ".",
        "resource",
        "trace/internal",
        "plugin/ocgrpc",
        "stats/view",
        "trace/propagation",
        "stats/internal",
        "metric/metricproducer",
        "trace",
        "metric/metricdata",
        "trace/tracestate",
        "plugin/ochttp",
        "stats",
        "internal/tagencoding",
        "tag",
        "internal",
        "plugin/ochttp/propagation/b3",
    ],
    deps = [
        ":google.golang.org.grpc",
        ":github.com.golang.groupcache",
    ],
)

go_module(
    name = "github.com.go-test.deep",
    module = "github.com/go-test/deep",
    version = "v1.0.7",
)

go_module(
    name = "github.com.gammazero.deque",
    module = "github.com/gammazero/deque",
    version = "v0.0.0-20190130191400-2afb3858e9c7",
)

go_module(
    name = "golang.org.x.net",
    module = "golang.org/x/net",
    version = "v0.0.0-20210520170846-37e1c6afe023",
    install = [
        "ipv6",
        "internal/socket",
        "bpf",
        "http2",
        "http2/hpack",
        "internal/iana",
        "trace",
        "ipv4",
        "idna",
        "context/ctxhttp",
        "http/httpproxy",
        "http/httpguts",
        "internal/timeseries",
        "publicsuffix",
        "context",
        "internal/socks",
        "proxy",
    ],
    deps = [
        ":golang.org.x.sys",
        ":golang.org.x.text",
    ],
)

go_module(
    name = "github.com.jcmturner.goidentity.v6",
    module = "github.com/jcmturner/goidentity/v6",
    version = "v6.0.1",
    deps = [
        ":github.com.hashicorp.go-uuid",
    ],
)

go_module(
    name = "github.com.cenkalti.backoff",
    module = "github.com/cenkalti/backoff",
    version = "v2.2.1+incompatible",
)

go_module(
    name = "github.com.NYTimes.gziphandler",
    module = "github.com/NYTimes/gziphandler",
    version = "v1.1.1",
)

go_module(
    name = "github.com.BurntSushi.toml",
    module = "github.com/BurntSushi/toml",
    version = "v0.3.1",
)

go_module(
    name = "github.com.hashicorp.go-multierror",
    module = "github.com/hashicorp/go-multierror",
    version = "v1.1.1",
    deps = [
        ":github.com.hashicorp.errwrap",
    ],
)

go_module(
    name = "github.com.prometheus.common",
    module = "github.com/prometheus/common",
    version = "v0.11.1",
    install = [
        "model",
        "expfmt",
        "internal/bitbucket.org/ww/goautoneg",
    ],
    deps = [
        ":github.com.matttproud.golang_protobuf_extensions",
        ":github.com.golang.protobuf",
        ":github.com.prometheus.client_model",
    ],
)

go_module(
    name = "github.com.beorn7.perks",
    module = "github.com/beorn7/perks",
    version = "v1.0.1",
    install = [
        "quantile",
    ],
)

go_module(
    name = "github.com.google.gofuzz",
    module = "github.com/google/gofuzz",
    version = "v1.1.0",
)

go_module(
    name = "github.com.hashicorp.go-kms-wrapping.entropy",
    module = "github.com/hashicorp/go-kms-wrapping/entropy",
    version = "v0.1.0",
)

go_module(
    name = "github.com.dimchansky.utfbom",
    module = "github.com/dimchansky/utfbom",
    version = "v1.1.1",
)

go_module(
    name = "github.com.golang.groupcache",
    module = "github.com/golang/groupcache",
    version = "v0.0.0-20210331224755-41bb18bfe9da",
    install = [
        "lru",
    ],
)

go_module(
    name = "github.com.softlayer.softlayer-go",
    module = "github.com/softlayer/softlayer-go",
    version = "v0.0.0-20180806151055-260589d94c7d",
    install = [
        "sl",
        "services",
        "session",
        "filter",
        "config",
        "datatypes",
    ],
    deps = [
        ":github.com.renier.xmlrpc",
    ],
)

go_module(
    name = "github.com.aliyun.alibaba-cloud-sdk-go",
    module = "github.com/aliyun/alibaba-cloud-sdk-go",
    version = "v0.0.0-20190620160927-9418d7b0cd0f",
    install = [
        "sdk/auth/credentials/providers",
        "services/sts",
        "sdk/endpoints",
        "sdk/requests",
        "sdk/auth/credentials",
        "sdk/auth/signers",
        "sdk/errors",
        "sdk",
        "services/ram",
        "sdk/auth/credentials/provider",
        "sdk/responses",
        "sdk/utils",
        "sdk/auth",
        "services/kms",
    ],
    deps = [
        ":github.com.jmespath.go-jmespath",
        ":gopkg.in.ini.v1",
        ":github.com.json-iterator.go",
    ],
)

go_module(
    name = "github.com.hashicorp.serf",
    module = "github.com/hashicorp/serf",
    version = "v0.9.5",
    install = [
        "coordinate",
    ],
    deps = [
        ":github.com.armon.go-metrics",
    ],
)

go_module(
    name = "github.com.hashicorp.vault-plugin-database-snowflake",
    module = "github.com/hashicorp/vault-plugin-database-snowflake",
    version = "v0.2.1",
    deps = [
        ":github.com.hashicorp.vault.sdk_1",
        ":github.com.hashicorp.errwrap",
        ":github.com.snowflakedb.gosnowflake",
        ":github.com.hashicorp.go-secure-stdlib.strutil",
    ],
)

go_module(
    name = "github.com.hashicorp.vault-plugin-secrets-terraform",
    module = "github.com/hashicorp/vault-plugin-secrets-terraform",
    version = "v0.2.0",
    deps = [
        ":github.com.hashicorp.vault.sdk_1",
        ":github.com.hashicorp.go-tfe",
    ],
)

go_module(
    name = "github.com.mattn.go-ieproxy",
    module = "github.com/mattn/go-ieproxy",
    version = "v0.0.1",
)

go_module(
    name = "gopkg.in.square.go-jose.v2",
    module = "gopkg.in/square/go-jose.v2",
    version = "v2.5.1",
    install = [
        ".",
        "json",
        "jwt",
        "cipher",
    ],
    deps = [
        ":golang.org.x.crypto",
    ],
)

go_module(
    name = "github.com.hashicorp.golang-lru",
    module = "github.com/hashicorp/golang-lru",
    version = "v0.5.4",
    install = [
        ".",
        "simplelru",
    ],
)

go_module(
    name = "github.com.aws.aws-sdk-go",
    module = "github.com/aws/aws-sdk-go",
    version = "v1.37.19",
    install = [
        "aws/credentials",
        "private/protocol/rest",
        "aws/csm",
        "service/sso/ssoiface",
        "service/dynamodb",
        "internal/s3shared/s3err",
        "private/protocol/json/jsonutil",
        "internal/sync/singleflight",
        "service/s3",
        "aws/defaults",
        "service/dynamodb/dynamodbattribute",
        "service/sts/stsiface",
        "aws/credentials/endpointcreds",
        "aws",
        "service/sts",
        "aws/client",
        "private/protocol/query",
        "aws/awsutil",
        "private/checksum",
        "private/protocol/query/queryutil",
        "internal/sdkio",
        "internal/strings",
        "private/protocol/restjson",
        "aws/credentials/processcreds",
        "aws/ec2metadata",
        "aws/awserr",
        "service/ec2",
        "aws/corehandlers",
        "internal/ini",
        "aws/request",
        "private/protocol/xml/xmlutil",
        "aws/crr",
        "private/protocol/eventstream/eventstreamapi",
        "aws/session",
        "aws/signer/v4",
        "internal/sdkrand",
        "internal/s3shared",
        "aws/endpoints",
        "aws/arn",
        "aws/client/metadata",
        "internal/shareddefaults",
        "aws/credentials/ssocreds",
        "private/protocol/ec2query",
        "private/protocol/eventstream",
        "service/iam/iamiface",
        "private/protocol/jsonrpc",
        "private/protocol",
        "private/protocol/restxml",
        "service/kms",
        "aws/credentials/ec2rolecreds",
        "internal/sdkuri",
        "service/iam",
        "service/sso",
        "service/kms/kmsiface",
        "aws/credentials/stscreds",
        "internal/s3shared/arn",
        "internal/sdkmath",
    ],
    deps = [
        ":github.com.jmespath.go-jmespath",
    ],
)

go_module(
    name = "github.com.hashicorp.vault-plugin-auth-kubernetes",
    module = "github.com/hashicorp/vault-plugin-auth-kubernetes",
    version = "v0.10.1",
    deps = [
        ":k8s.io.apimachinery",
        ":github.com.briankassouf.jose",
        ":github.com.hashicorp.go-multierror",
        ":github.com.hashicorp.go-sockaddr",
        ":github.com.hashicorp.vault.sdk_1",
        ":github.com.hashicorp.errwrap",
        ":k8s.io.api",
        ":github.com.mitchellh.mapstructure",
        ":github.com.hashicorp.go-cleanhttp",
    ],
)

go_module(
    name = "github.com.hashicorp.vault-plugin-database-couchbase",
    module = "github.com/hashicorp/vault-plugin-database-couchbase",
    version = "v0.4.1",
    deps = [
        ":github.com.hashicorp.vault.sdk_1",
        ":github.com.mitchellh.mapstructure",
        ":github.com.hashicorp.go-hclog",
        ":github.com.couchbase.gocb.v2",
        ":github.com.hashicorp.errwrap",
        ":github.com.hashicorp.go-version",
        ":github.com.cenkalti.backoff",
    ],
)

go_module(
    name = "github.com.containerd.containerd",
    module = "github.com/containerd/containerd",
    version = "v1.4.3",
    install = [
        "errdefs",
    ],
    deps = [
        ":github.com.pkg.errors",
        ":google.golang.org.grpc",
    ],
)

go_module(
    name = "github.com.Azure.go-autorest.logger",
    module = "github.com/Azure/go-autorest/logger",
    version = "v0.2.1",
)

go_module(
    name = "github.com.mitchellh.go-wordwrap",
    module = "github.com/mitchellh/go-wordwrap",
    version = "v1.0.0",
)

go_module(
    name = "github.com.google.go-cmp",
    module = "github.com/google/go-cmp",
    version = "v0.5.5",
    install = [
        "cmp",
        "cmp/internal/flags",
        "cmp/internal/value",
        "cmp/internal/diff",
        "cmp/internal/function",
    ],
)

go_module(
    name = "github.com.yuin.gopher-lua",
    module = "github.com/yuin/gopher-lua",
    version = "v0.0.0-20200816102855-ee81675732da",
    install = [
        ".",
        "ast",
        "parse",
        "pm",
    ],
)

go_module(
    name = "github.com.hashicorp.go-msgpack",
    module = "github.com/hashicorp/go-msgpack",
    version = "v1.1.5",
    install = [
        "codec",
    ],
)

go_module(
    name = "github.com.hashicorp.go-version",
    module = "github.com/hashicorp/go-version",
    version = "v1.2.1",
)

go_module(
    name = "github.com.jefferai.isbadcipher",
    module = "github.com/jefferai/isbadcipher",
    version = "v0.0.0-20190226160619-51d2077c035f",
)

go_module(
    name = "google.golang.org.grpc",
    module = "google.golang.org/grpc",
    version = "v1.29.1",
    install = [
        ".",
        "credentials/alts/internal/handshaker",
        "encoding/proto",
        "resolver",
        "balancer/base",
        "status",
        "metadata",
        "internal/grpcsync",
        "keepalive",
        "internal/backoff",
        "internal/resolver/dns",
        "credentials/alts/internal/proto/grpc_gcp",
        "balancer",
        "credentials",
        "balancer/grpclb/grpc_lb_v1",
        "internal/grpclog",
        "peer",
        "internal",
        "connectivity",
        "grpclog",
        "internal/resolver/passthrough",
        "tap",
        "internal/binarylog",
        "health/grpc_health_v1",
        "internal/balancerload",
        "credentials/alts/internal",
        "binarylog/grpc_binarylog_v1",
        "internal/syscall",
        "internal/status",
        "internal/grpcutil",
        "internal/channelz",
        "credentials/alts/internal/authinfo",
        "codes",
        "credentials/google",
        "internal/grpcrand",
        "credentials/alts/internal/handshaker/service",
        "credentials/internal",
        "credentials/oauth",
        "stats",
        "attributes",
        "internal/envconfig",
        "credentials/alts",
        "balancer/roundrobin",
        "balancer/grpclb",
        "naming",
        "internal/transport",
        "resolver/passthrough",
        "backoff",
        "resolver/dns",
        "encoding",
        "serviceconfig",
        "internal/buffer",
        "credentials/alts/internal/conn",
        "health",
    ],
    deps = [
        ":github.com.golang.protobuf",
        ":google.golang.org.genproto_1",
        ":golang.org.x.net",
        ":golang.org.x.sys",
        ":golang.org.x.oauth2",
        ":github.com.google.go-cmp",
    ],
)

go_module(
    name = "github.com.go-ldap.ldap.v3",
    module = "github.com/go-ldap/ldap/v3",
    version = "v3.2.4",
    deps = [
        ":github.com.Azure.go-ntlmssp",
        ":github.com.go-asn1-ber.asn1-ber",
    ],
)

go_module(
    name = "github.com.prometheus.procfs",
    module = "github.com/prometheus/procfs",
    version = "v0.1.3",
    install = [
        ".",
        "internal/util",
        "internal/fs",
    ],
    deps = [
        ":golang.org.x.sys",
    ],
)

go_module(
    name = "github.com.Azure.go-autorest.autorest.adal",
    module = "github.com/Azure/go-autorest/autorest/adal",
    version = "v0.9.13",
    deps = [
        ":github.com.form3tech-oss.jwt-go",
        ":github.com.Azure.go-autorest.logger",
        ":github.com.Azure.go-autorest.autorest.date",
        ":github.com.Azure.go-autorest.tracing",
        ":golang.org.x.crypto",
    ],
)

go_module(
    name = "github.com.circonus-labs.circonus-gometrics",
    module = "github.com/circonus-labs/circonus-gometrics",
    version = "v2.3.1+incompatible",
    install = [
        ".",
        "api",
        "api/config",
        "checkmgr",
    ],
    deps = [
        ":github.com.hashicorp.go-retryablehttp",
        ":github.com.pkg.errors",
        ":github.com.tv42.httpunix",
        ":github.com.circonus-labs.circonusllhist",
    ],
)

go_module(
    name = "cloud.google.com.go.spanner",
    module = "cloud.google.com/go/spanner",
    version = "v1.5.1",
    install = [
        ".",
        "admin/instance/apiv1",
        "apiv1",
    ],
    deps = [
        ":google.golang.org.grpc",
        ":cloud.google.com.go",
        ":google.golang.org.genproto",
        ":google.golang.org.api",
        ":github.com.googleapis.gax-go.v2",
        ":github.com.golang.protobuf",
        ":go.opencensus.io",
        ":google.golang.org.genproto_1",
        ":cloud.google.com.go_1",
    ],
)

go_module(
    name = "github.com.dustin.go-humanize",
    module = "github.com/dustin/go-humanize",
    version = "v1.0.0",
)

go_module(
    name = "golang.org.x.term",
    module = "golang.org/x/term",
    version = "v0.0.0-20210220032956-6a3ed077a48d",
    deps = [
        ":golang.org.x.sys",
    ],
)

go_module(
    name = "github.com.matttproud.golang_protobuf_extensions",
    module = "github.com/matttproud/golang_protobuf_extensions",
    version = "v1.0.1",
    install = [
        "pbutil",
    ],
    deps = [
        ":github.com.golang.protobuf",
    ],
)

go_module(
    name = "github.com.modern-go.concurrent",
    module = "github.com/modern-go/concurrent",
    version = "v0.0.0-20180306012644-bacd9c7ef1dd",
)

go_module(
    name = "github.com.jcmturner.gofork",
    module = "github.com/jcmturner/gofork",
    version = "v1.0.0",
    install = [
        "encoding/asn1",
        "x/crypto/pbkdf2",
    ],
)

go_module(
    name = "github.com.aerospike.aerospike-client-go",
    module = "github.com/aerospike/aerospike-client-go",
    version = "v3.1.1+incompatible",
    install = [
        ".",
        "pkg/ripemd160",
        "utils/buffer",
        "internal/lua",
        "types",
        "internal/lua/resources",
        "internal/particle_type",
        "types/rand",
        "internal/atomic",
        "logger",
        "pkg/bcrypt",
    ],
    deps = [
        ":github.com.yuin.gopher-lua",
        ":golang.org.x.sync",
    ],
)

go_module(
    name = "github.com.Azure.go-autorest.autorest.date",
    module = "github.com/Azure/go-autorest/autorest/date",
    version = "v0.3.0",
)

go_module(
    name = "github.com.patrickmn.go-cache",
    module = "github.com/patrickmn/go-cache",
    version = "v2.1.0+incompatible",
)

go_module(
    name = "github.com.coreos.go-systemd.v22",
    module = "github.com/coreos/go-systemd/v22",
    version = "v22.0.0",
    install = [
        "journal",
    ],
)

go_module(
    name = "github.com.cloudfoundry-community.go-cfclient",
    module = "github.com/cloudfoundry-community/go-cfclient",
    version = "v0.0.0-20190201205600-f136f9222381",
    deps = [
        ":golang.org.x.oauth2",
        ":github.com.pkg.errors",
        ":github.com.Masterminds.semver",
        ":code.cloudfoundry.org.gofileutils",
        ":golang.org.x.net",
    ],
)

go_module(
    name = "github.com.hashicorp.go-raftchunking",
    module = "github.com/hashicorp/go-raftchunking",
    version = "v0.6.3-0.20191002164813-7e9e8525653a",
    install = [
        ".",
        "types",
    ],
    deps = [
        ":github.com.hashicorp.raft",
        ":github.com.hashicorp.errwrap",
        ":github.com.mitchellh.copystructure",
        ":github.com.golang.protobuf",
    ],
)

go_module(
    name = "github.com.lib.pq",
    module = "github.com/lib/pq",
    version = "v1.8.0",
    install = [
        ".",
        "oid",
        "scram",
    ],
)

go_module(
    name = "github.com.hashicorp.vault-plugin-secrets-openldap",
    module = "github.com/hashicorp/vault-plugin-secrets-openldap",
    version = "v0.5.1",
    install = [
        ".",
        "client",
    ],
    deps = [
        ":github.com.go-ldap.ldap.v3",
        ":github.com.hashicorp.vault.sdk_1",
        ":github.com.mitchellh.mapstructure",
        ":github.com.go-ldap.ldif",
        ":github.com.hashicorp.go-multierror",
        ":github.com.hashicorp.errwrap",
        ":golang.org.x.text",
    ],
)

go_module(
    name = "github.com.couchbase.gocb.v2",
    module = "github.com/couchbase/gocb/v2",
    version = "v2.1.4",
    install = [
        ".",
        "search",
    ],
    deps = [
        ":github.com.couchbase.gocbcore.v9",
        ":github.com.pkg.errors",
        ":github.com.google.uuid",
    ],
)

go_module(
    name = "github.com.google.uuid",
    module = "github.com/google/uuid",
    version = "v1.1.2",
)

go_module(
    name = "layeh.com.radius",
    module = "layeh.com/radius",
    version = "v0.0.0-20190322222518-890bc1058917",
    install = [
        ".",
        "rfc2865",
    ],
)

go_module(
    name = "github.com.ulikunitz.xz",
    module = "github.com/ulikunitz/xz",
    version = "v0.5.7",
    install = [
        ".",
        "internal/hash",
        "lzma",
        "internal/xlog",
    ],
)

go_module(
    name = "github.com.hashicorp.go-slug",
    module = "github.com/hashicorp/go-slug",
    version = "v0.4.1",
)

go_module(
    name = "k8s.io.klog.v2",
    module = "k8s.io/klog/v2",
    version = "v2.9.0",
    deps = [
        ":github.com.go-logr.logr",
    ],
)

go_module(
    name = "github.com.hashicorp.vault-plugin-secrets-azure",
    module = "github.com/hashicorp/vault-plugin-secrets-azure",
    version = "v0.10.0",
    deps = [
        ":github.com.hashicorp.vault.sdk_1",
        ":github.com.Azure.go-autorest.autorest.to",
        ":github.com.Azure.azure-sdk-for-go",
        ":github.com.Azure.go-autorest.autorest.date",
        ":github.com.mitchellh.mapstructure",
        ":github.com.Azure.go-autorest.autorest",
        ":github.com.hashicorp.errwrap",
        ":github.com.Azure.go-autorest.autorest.azure.auth",
        ":github.com.hashicorp.go-uuid",
        ":github.com.hashicorp.go-multierror",
    ],
)

go_module(
    name = "github.com.hashicorp.vault-plugin-auth-azure",
    module = "github.com/hashicorp/vault-plugin-auth-azure",
    version = "v0.8.0",
    deps = [
        ":github.com.hashicorp.errwrap",
        ":github.com.Azure.go-autorest.autorest",
        ":github.com.hashicorp.go-cleanhttp",
        ":golang.org.x.oauth2",
        ":github.com.hashicorp.vault.sdk_1",
        ":github.com.Azure.go-autorest.autorest.azure.auth",
        ":github.com.Azure.azure-sdk-for-go",
        ":github.com.coreos.go-oidc",
        ":github.com.Azure.go-autorest.autorest.to",
    ],
)

go_module(
    name = "github.com.armon.go-proxyproto",
    module = "github.com/armon/go-proxyproto",
    version = "v0.0.0-20210323213023-7e956b284f0a",
)

go_module(
    name = "github.com.mattn.go-colorable",
    module = "github.com/mattn/go-colorable",
    version = "v0.1.8",
    deps = [
        ":github.com.mattn.go-isatty",
    ],
)

go_module(
    name = "golang.org.x.crypto",
    module = "golang.org/x/crypto",
    version = "v0.0.0-20210513164829-c07d793c2f9a",
    install = [
        "scrypt",
        "cryptobyte/asn1",
        "chacha20",
        "poly1305",
        "pkcs12/internal/rc2",
        "ed25519",
        "ocsp",
        "bcrypt",
        "md4",
        "blake2b",
        "curve25519",
        "ssh/agent",
        "ssh/internal/bcrypt_pbkdf",
        "pbkdf2",
        "cryptobyte",
        "curve25519/internal/field",
        "hkdf",
        "ssh",
        "blowfish",
        "pkcs12",
        "chacha20poly1305",
        "internal/subtle",
        "ssh/terminal",
    ],
    deps = [
        ":golang.org.x.sys",
        ":golang.org.x.term",
    ],
)

go_module(
    name = "github.com.jefferai.jsonx",
    module = "github.com/jefferai/jsonx",
    version = "v1.0.0",
    deps = [
        ":github.com.Jeffail.gabs",
    ],
)

go_module(
    name = "github.com.hashicorp.vault-plugin-secrets-alicloud",
    module = "github.com/hashicorp/vault-plugin-secrets-alicloud",
    version = "v0.9.0",
    install = [
        ".",
        "clients",
    ],
    deps = [
        ":github.com.hashicorp.vault.sdk_1",
        ":github.com.aliyun.alibaba-cloud-sdk-go",
        ":github.com.hashicorp.go-multierror",
        ":github.com.hashicorp.go-uuid",
    ],
)

go_module(
    name = "github.com.docker.go-connections",
    module = "github.com/docker/go-connections",
    version = "v0.4.0",
    install = [
        "nat",
        "sockets",
        "tlsconfig",
    ],
    deps = [
        ":golang.org.x.net",
        ":github.com.pkg.errors",
    ],
)

go_module(
    name = "github.com.docker.distribution",
    module = "github.com/docker/distribution",
    version = "v2.7.1+incompatible",
    install = [
        "reference",
        "digestset",
        "registry/api/errcode",
    ],
    deps = [
        ":github.com.opencontainers.go-digest",
    ],
)

go_module(
    name = "github.com.moby.spdystream",
    module = "github.com/moby/spdystream",
    version = "v0.2.0",
    install = [
        ".",
        "spdy",
    ],
)

go_module(
    name = "k8s.io.client-go",
    module = "k8s.io/client-go",
    version = "v0.22.1",
    install = [
        "tools/auth",
        "transport/spdy",
        "kubernetes/typed/apps/v1",
        "listers/apps/v1beta1",
        "third_party/forked/golang/template",
        "kubernetes/typed/storage/v1alpha1",
        "util/retry",
        "listers/extensions/v1beta1",
        "applyconfigurations/batch/v1beta1",
        "listers/rbac/v1",
        "informers/scheduling/v1alpha1",
        "kubernetes/typed/storage/v1/fake",
        "kubernetes/typed/admissionregistration/v1",
        "kubernetes/typed/scheduling/v1alpha1",
        "tools/clientcmd/api/latest",
        "kubernetes/typed/node/v1alpha1",
        "kubernetes/typed/autoscaling/v1",
        "applyconfigurations/internal",
        "plugin/pkg/client/auth/oidc",
        "applyconfigurations/events/v1beta1",
        "dynamic/dynamiclister",
        "listers/flowcontrol/v1alpha1",
        "applyconfigurations/extensions/v1beta1",
        "kubernetes/typed/flowcontrol/v1beta1/fake",
        "kubernetes/typed/admissionregistration/v1/fake",
        "listers/coordination/v1",
        "applyconfigurations/imagepolicy/v1alpha1",
        "informers/apps/v1beta2",
        "dynamic/fake",
        "informers/coordination/v1",
        "informers/admissionregistration/v1",
        "tools/cache/testing",
        "kubernetes/typed/node/v1beta1/fake",
        "listers/admissionregistration/v1beta1",
        "kubernetes/typed/apps/v1/fake",
        "informers/node",
        "kubernetes/typed/batch/v1beta1/fake",
        "transport",
        "applyconfigurations/networking/v1",
        "informers/apps",
        "informers/discovery/v1beta1",
        "kubernetes/typed/certificates/v1beta1/fake",
        "tools/remotecommand",
        "informers/events/v1beta1",
        "informers/policy",
        "informers/core",
        "applyconfigurations/certificates/v1beta1",
        "plugin/pkg/client/auth/exec",
        "plugin/pkg/client/auth",
        "scale",
        "kubernetes/typed/certificates/v1beta1",
        "listers/apps/v1",
        "applyconfigurations/scheduling/v1",
        "informers/certificates/v1",
        "kubernetes/typed/apps/v1beta1/fake",
        "applyconfigurations/node/v1",
        "kubernetes/typed/storage/v1beta1/fake",
        "pkg/apis/clientauthentication",
        "kubernetes/typed/scheduling/v1",
        "informers/networking",
        "kubernetes/typed/flowcontrol/v1alpha1/fake",
        "applyconfigurations/networking/v1beta1",
        "kubernetes/typed/policy/v1beta1/fake",
        "discovery/cached",
        "informers/batch",
        "dynamic/dynamicinformer",
        "tools/auth/exec",
        "informers/apiserverinternal",
        "applyconfigurations/flowcontrol/v1alpha1",
        "informers/autoscaling/v2beta1",
        "scale/scheme/extensionsint",
        "util/cert",
        "kubernetes/typed/apiserverinternal/v1alpha1",
        "listers/admissionregistration/v1",
        "kubernetes/typed/scheduling/v1beta1",
        "informers/extensions/v1beta1",
        "informers/node/v1",
        "util/workqueue",
        "listers/node/v1",
        "applyconfigurations/events/v1",
        "kubernetes/typed/authorization/v1beta1",
        "kubernetes/typed/node/v1beta1",
        "tools/reference",
        "listers/authorization/v1beta1",
        "listers/storage/v1",
        "pkg/version",
        "kubernetes/typed/certificates/v1/fake",
        "applyconfigurations/meta/v1",
        "informers/coordination",
        "applyconfigurations/apps/v1",
        "util/flowcontrol",
        "informers/storage/v1alpha1",
        "tools/watch",
        "informers/admissionregistration",
        "applyconfigurations/coordination/v1beta1",
        "listers/rbac/v1alpha1",
        "scale/scheme/extensionsv1beta1",
        "tools/metrics",
        "listers/autoscaling/v2beta2",
        "informers/policy/v1",
        "informers/networking/v1",
        "scale/scheme/appsv1beta2",
        "kubernetes/typed/policy/v1/fake",
        "listers/batch/v1beta1",
        "kubernetes/typed/authorization/v1beta1/fake",
        "applyconfigurations/certificates/v1",
        "listers/networking/v1beta1",
        "kubernetes/typed/storage/v1",
        "kubernetes/typed/authentication/v1beta1",
        "kubernetes/typed/apiserverinternal/v1alpha1/fake",
        "informers/storage",
        "kubernetes/typed/rbac/v1/fake",
        "listers/node/v1beta1",
        "tools/clientcmd",
        "util/connrotation",
        "tools/leaderelection",
        "kubernetes/fake",
        "listers/networking/v1",
        "informers/autoscaling/v1",
        "kubernetes/typed/core/v1",
        "informers/apiserverinternal/v1alpha1",
        "applyconfigurations/rbac/v1beta1",
        "informers/scheduling",
        "applyconfigurations/core/v1",
        "informers/flowcontrol/v1alpha1",
        "informers/batch/v1beta1",
        "listers/scheduling/v1",
        "tools/clientcmd/api",
        "applyconfigurations/autoscaling/v2beta2",
        "kubernetes/typed/batch/v1",
        "kubernetes/typed/coordination/v1/fake",
        "kubernetes/typed/discovery/v1/fake",
        "metadata/metadatalister",
        "discovery/cached/memory",
        "kubernetes/typed/authentication/v1beta1/fake",
        "applyconfigurations/apps/v1beta2",
        "kubernetes/typed/rbac/v1beta1",
        "dynamic",
        "informers/core/v1",
        "listers/imagepolicy/v1alpha1",
        "listers/autoscaling/v2beta1",
        "listers/batch/v1",
        "plugin/pkg/client/auth/gcp",
        "applyconfigurations/autoscaling/v2beta1",
        "kubernetes/typed/autoscaling/v1/fake",
        "kubernetes/typed/networking/v1beta1",
        "informers/scheduling/v1beta1",
        "applyconfigurations/storage/v1beta1",
        "informers/rbac/v1alpha1",
        "informers/events",
        "applyconfigurations/scheduling/v1beta1",
        "kubernetes/typed/discovery/v1",
        "examples/dynamic-create-update-delete-deployment",
        "restmapper",
        "kubernetes/typed/policy/v1beta1",
        "informers/policy/v1beta1",
        "listers/apps/v1beta2",
        "applyconfigurations/admissionregistration/v1",
        "listers/autoscaling/v1",
        "rest/fake",
        "kubernetes/typed/autoscaling/v2beta2",
        "applyconfigurations/coordination/v1",
        "util/certificate",
        "pkg/apis/clientauthentication/install",
        "kubernetes/typed/scheduling/v1alpha1/fake",
        "kubernetes/typed/apps/v1beta2/fake",
        "tools/record",
        "applyconfigurations",
        "kubernetes/typed/networking/v1",
        "pkg/apis/clientauthentication/v1beta1",
        "examples/leader-election",
        "informers/discovery/v1",
        "informers/admissionregistration/v1beta1",
        "kubernetes/typed/certificates/v1",
        "kubernetes",
        "listers/flowcontrol/v1beta1",
        "listers/apiserverinternal/v1alpha1",
        "kubernetes/typed/storage/v1alpha1/fake",
        "listers/node/v1alpha1",
        "kubernetes/typed/apps/v1beta1",
        "util/homedir",
        "informers/scheduling/v1",
        "listers/discovery/v1",
        "tools/leaderelection/resourcelock",
        "rest",
        "informers/autoscaling",
        "examples/out-of-cluster-client-configuration",
        "util/jsonpath",
        "discovery/fake",
        "applyconfigurations/policy/v1beta1",
        "kubernetes/typed/events/v1beta1",
        "informers/autoscaling/v2beta2",
        "listers/certificates/v1beta1",
        "kubernetes/scheme",
        "listers/core/v1",
        "kubernetes/typed/autoscaling/v2beta1",
        "applyconfigurations/discovery/v1beta1",
        "informers/node/v1beta1",
        "util/keyutil",
        "discovery",
        "plugin/pkg/client/auth/openstack",
        "metadata",
        "kubernetes/typed/node/v1",
        "listers/authentication/v1",
        "informers/storage/v1beta1",
        "util/exec",
        "kubernetes/typed/rbac/v1alpha1/fake",
        "kubernetes/typed/flowcontrol/v1beta1",
        "rest/watch",
        "kubernetes/typed/networking/v1/fake",
        "applyconfigurations/flowcontrol/v1beta1",
        "applyconfigurations/batch/v1",
        "listers/authentication/v1beta1",
        "kubernetes/typed/authentication/v1/fake",
        "informers/networking/v1beta1",
        "metadata/fake",
        "tools/record/util",
        "applyconfigurations/autoscaling/v1",
        "kubernetes/typed/authentication/v1",
        "scale/scheme",
        "tools/cache",
        "kubernetes/typed/rbac/v1",
        "informers/rbac",
        "listers/scheduling/v1alpha1",
        "applyconfigurations/storage/v1",
        "tools/events",
        "informers/discovery",
        "applyconfigurations/node/v1alpha1",
        "kubernetes/typed/extensions/v1beta1/fake",
        "kubernetes/typed/storage/v1beta1",
        "tools/pager",
        "listers/events/v1beta1",
        "pkg/apis/clientauthentication/v1",
        "kubernetes/typed/apps/v1beta2",
        "plugin/pkg/client/auth/azure",
        "tools/portforward",
        "kubernetes/typed/scheduling/v1beta1/fake",
        "examples/workqueue",
        "kubernetes/typed/flowcontrol/v1alpha1",
        "pkg/apis/clientauthentication/v1alpha1",
        "informers/coordination/v1beta1",
        "kubernetes/typed/policy/v1",
        "scale/scheme/autoscalingv1",
        "informers/storage/v1",
        "kubernetes/typed/coordination/v1beta1/fake",
        "kubernetes/typed/node/v1/fake",
        "kubernetes/typed/batch/v1/fake",
        "informers/apps/v1",
        "informers/certificates/v1beta1",
        "kubernetes/typed/autoscaling/v2beta2/fake",
        "examples/in-cluster-client-configuration",
        "metadata/metadatainformer",
        "listers/policy/v1beta1",
        "listers/rbac/v1beta1",
        "applyconfigurations/rbac/v1",
        "examples/create-update-delete-deployment",
        "applyconfigurations/scheduling/v1alpha1",
        "listers/events/v1",
        "informers/batch/v1",
        "applyconfigurations/apiserverinternal/v1alpha1",
        "informers/apps/v1beta1",
        "kubernetes/typed/rbac/v1alpha1",
        "informers/events/v1",
        "listers/storage/v1beta1",
        "scale/scheme/appsv1beta1",
        "informers/internalinterfaces",
        "scale/fake",
        "scale/scheme/appsint",
        "informers/extensions",
        "listers/storage/v1alpha1",
        "kubernetes/typed/rbac/v1beta1/fake",
        "listers/certificates/v1",
        "kubernetes/typed/scheduling/v1/fake",
        "kubernetes/typed/batch/v1beta1",
        "informers/rbac/v1",
        "kubernetes/typed/admissionregistration/v1beta1",
        "kubernetes/typed/node/v1alpha1/fake",
        "informers/rbac/v1beta1",
        "kubernetes/typed/events/v1/fake",
        "informers/node/v1alpha1",
        "listers/scheduling/v1beta1",
        "kubernetes/typed/extensions/v1beta1",
        "applyconfigurations/node/v1beta1",
        "kubernetes/typed/discovery/v1beta1",
        "informers/flowcontrol",
        "kubernetes/typed/authorization/v1",
        "applyconfigurations/policy/v1",
        "kubernetes/typed/coordination/v1",
        "listers/coordination/v1beta1",
        "listers/policy/v1",
        "applyconfigurations/admissionregistration/v1beta1",
        "informers",
        "listers/discovery/v1beta1",
        "discovery/cached/disk",
        "kubernetes/typed/admissionregistration/v1beta1/fake",
        "kubernetes/typed/events/v1",
        "applyconfigurations/rbac/v1alpha1",
        "util/certificate/csr",
        "tools/clientcmd/api/v1",
        "applyconfigurations/discovery/v1",
        "util/testing",
        "kubernetes/typed/core/v1/fake",
        "informers/flowcontrol/v1beta1",
        "listers/authorization/v1",
        "kubernetes/typed/networking/v1beta1/fake",
        "kubernetes/typed/events/v1beta1/fake",
        "testing",
        "kubernetes/typed/authorization/v1/fake",
        "kubernetes/typed/discovery/v1beta1/fake",
        "kubernetes/typed/coordination/v1beta1",
        "kubernetes/typed/autoscaling/v2beta1/fake",
        "informers/certificates",
        "examples/fake-client",
        "applyconfigurations/apps/v1beta1",
        "applyconfigurations/storage/v1alpha1",
    ],
    deps = [
        ":k8s.io.apimachinery",
        ":k8s.io.api",
        ":sigs.k8s.io.structured-merge-diff.v4",
        ":k8s.io.klog.v2",
        ":golang.org.x.oauth2",
        ":golang.org.x.term",
        ":github.com.davecgh.go-spew",
        ":golang.org.x.time",
        ":k8s.io.utils",
        ":github.com.spf13.pflag",
        ":github.com.imdario.mergo",
        ":github.com.googleapis.gnostic",
        ":github.com.golang.groupcache",
        ":github.com.google.uuid",
        ":golang.org.x.net",
        ":github.com.golang.protobuf",
        ":github.com.Azure.go-autorest.autorest",
        ":github.com.Azure.go-autorest.autorest.adal",
        ":github.com.gregjones.httpcache",
        ":github.com.peterbourgon.diskv",
        ":github.com.evanphx.json-patch",
    ],
)

go_module(
    name = "github.com.tencentcloud.tencentcloud-sdk-go",
    module = "github.com/tencentcloud/tencentcloud-sdk-go",
    version = "v3.0.171+incompatible",
    install = [
        "tencentcloud/common",
        "tencentcloud/common/profile",
        "tencentcloud/common/errors",
        "tencentcloud/common/http",
        "tencentcloud/cvm/v20170312",
    ],
)

go_module(
    name = "github.com.hashicorp.go-cleanhttp",
    module = "github.com/hashicorp/go-cleanhttp",
    version = "v0.5.2",
)

go_module(
    name = "github.com.kr.text",
    module = "github.com/kr/text",
    version = "v0.2.0",
)

go_module(
    name = "github.com.circonus-labs.circonusllhist",
    module = "github.com/circonus-labs/circonusllhist",
    version = "v0.1.3",
)

go_module(
    name = "github.com.google.btree",
    module = "github.com/google/btree",
    version = "v1.0.1",
)

go_module(
    name = "github.com.samuel.go-zookeeper",
    module = "github.com/samuel/go-zookeeper",
    version = "v0.0.0-20190923202752-2cc03de413da",
    install = [
        "zk",
    ],
)

go_module(
    name = "github.com.shirou.gopsutil",
    module = "github.com/shirou/gopsutil",
    version = "v3.21.5+incompatible",
    install = [
        "cpu",
        "host",
        "internal/common",
        "disk",
        "mem",
    ],
    deps = [
        ":github.com.tklauser.go-sysconf",
        ":golang.org.x.sys",
    ],
)

go_module(
    name = "github.com.huandu.xstrings",
    module = "github.com/huandu/xstrings",
    version = "v1.3.2",
)

go_module(
    name = "github.com.form3tech-oss.jwt-go",
    module = "github.com/form3tech-oss/jwt-go",
    version = "v3.2.3+incompatible",
)

go_module(
    name = "github.com.aws.aws-sdk-go-v2.service.internal.accept-encoding",
    module = "github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding",
    version = "v1.0.4",
    deps = [
        ":github.com.aws.smithy-go",
    ],
)

go_module(
    name = "github.com.cespare.xxhash.v2",
    module = "github.com/cespare/xxhash/v2",
    version = "v2.1.1",
)

go_module(
    name = "code.cloudfoundry.org.gofileutils",
    module = "code.cloudfoundry.org/gofileutils",
    version = "v0.0.0-20170111115228-4d0c80011a0f",
    install = [
        "fileutils",
    ],
)

go_module(
    name = "github.com.hashicorp.vault",
    module = "github.com/hashicorp/vault",
    version = "v1.8.2",
    install = [
        ".",
        "helper/identity",
        "command/agent/auth/approle",
        "builtin/credential/github",
        "builtin/credential/cert/cmd/cert",
        "builtin/credential/aws/cmd/aws",
        "helper/mfa",
        "plugins/database/influxdb",
        "vault/diagnose",
        "command/agent/config",
        "builtin/logical/mongodb",
        "command/agent/cache/cachememdb",
        "command/agent/sink/mock",
        "internalshared/reloadutil",
        "plugins/database/mongodb/mongodb-database-plugin",
        "builtin/logical/totp",
        "builtin/logical/consul/cmd/consul",
        "internalshared/listenerutil",
        "serviceregistration/consul",
        "physical/postgresql",
        "builtin/credential/aws/pkcs7",
        "physical/etcd",
        "vault",
        "helper/dhutil",
        "helper/testhelpers/certhelpers",
        "helper/testhelpers/teststorage",
        "command/config",
        "command/agent/auth/aws",
        "internalshared/kv-builder",
        "helper/forwarding",
        "helper/namespace",
        "helper/testhelpers/cassandra",
        "physical/couchdb",
        "command/token",
        "serviceregistration",
        "internalshared/gatedwriter",
        "physical/alicloudoss",
        "builtin/logical/consul",
        "builtin/logical/postgresql",
        "builtin/logical/aws",
        "physical/manta",
        "serviceregistration/kubernetes",
        "vault/seal",
        "helper/pgpkeys",
        "helper/builtinplugins",
        "helper/identity/mfa",
        "builtin/credential/approle/cmd/approle",
        "builtin/audit/socket",
        "physical/cockroachdb",
        "helper/testhelpers/minio",
        "builtin/credential/radius",
        "builtin/logical/aws/cmd/aws",
        "helper/testhelpers/logical",
        "helper/testhelpers/ldap",
        "plugins/database/postgresql",
        "helper/monitor",
        "audit",
        "builtin/logical/pki/cmd/pki",
        "helper/testhelpers/docker",
        "command/agent/auth",
        "builtin/credential/ldap/cmd/ldap",
        "physical/oci",
        "builtin/logical/mysql",
        "builtin/logical/nomad",
        "plugins/database/hana/hana-database-plugin",
        "command/agent/auth/gcp",
        "command/agent/cache/cacheboltdb",
        "builtin/credential/github/cmd/github",
        "helper/xor",
        "helper/random",
        "builtin/logical/transit/cmd/transit",
        "helper/testhelpers/mongodb",
        "builtin/credential/userpass",
        "builtin/logical/pki",
        "command/server",
        "builtin/credential/okta",
        "command/agent",
        "plugins/database/mysql/mysql-database-plugin",
        "command/agent/cache",
        "builtin/logical/mssql/cmd/mssql",
        "shamir",
        "command/agent/auth/cert",
        "builtin/credential/ldap",
        "builtin/audit/syslog",
        "builtin/logical/mssql",
        "builtin/logical/database",
        "command/agent/cache/keymanager",
        "physical/dynamodb",
        "helper/flag-kv",
        "plugins/database/redshift",
        "physical/mysql",
        "physical/gcs",
        "helper/constants",
        "command/agent/auth/cf",
        "builtin/credential/app-id",
        "serviceregistration/kubernetes/client/cmd/kubeclient",
        "helper/timeutil",
        "vault/cluster",
        "builtin/credential/userpass/cmd/userpass",
        "builtin/audit/file",
        "plugins/database/influxdb/influxdb-database-plugin",
        "builtin/logical/rabbitmq/cmd/rabbitmq",
        "helper/hostutil",
        "physical/cassandra",
        "builtin/logical/mysql/cmd/mysql",
        "command/agent/winsvc",
        "serviceregistration/kubernetes/testing",
        "plugins/database/mssql",
        "command/agent/auth/kubernetes",
        "builtin/logical/cassandra/cmd/cassandra",
        "helper/testhelpers/teststorage/consul",
        "command/agent/auth/azure",
        "builtin/credential/token",
        "helper/testhelpers/postgresql",
        "helper/testhelpers/seal",
        "builtin/credential/cert",
        "plugins/database/cassandra/cassandra-database-plugin",
        "vault/replication",
        "builtin/credential/okta/cmd/okta",
        "vault/external_tests/sealmigration",
        "helper/testhelpers/azurite",
        "helper/fairshare",
        "helper/testhelpers/mysql",
        "helper/flag-slice",
        "plugins/database/mysql/mysql-legacy-database-plugin",
        "builtin/plugin",
        "physical/foundationdb",
        "plugins/database/postgresql/postgresql-database-plugin",
        "plugins/database/hana",
        "physical/raft",
        "builtin/logical/mongodb/cmd/mongodb",
        "command/agent/sink",
        "vault/quotas",
        "vault/activity",
        "physical/spanner",
        "internalshared/configutil",
        "plugins/database/mssql/mssql-database-plugin",
        "helper/mfa/duo",
        "physical/consul",
        "plugins/database/redshift/redshift-database-plugin",
        "helper/storagepacker",
        "builtin/credential/radius/cmd/radius",
        "command/agent/sink/inmem",
        "physical/aerospike",
        "helper/testhelpers/consul",
        "command/agent/sink/file",
        "helper/proxyutil",
        "builtin/logical/nomad/cmd/nomad",
        "physical/swift",
        "command/agent/auth/kerberos",
        "command",
        "builtin/logical/totp/cmd/totp",
        "command/agent/auth/jwt",
        "helper/policies",
        "builtin/logical/transit",
        "physical/mssql",
        "builtin/logical/ssh/cmd/ssh",
        "physical/azure",
        "physical/zookeeper",
        "helper/metricsutil",
        "builtin/logical/cassandra",
        "helper/testhelpers/mssql",
        "builtin/credential/app-id/cmd/app-id",
        "helper/testhelpers/fakegcsserver",
        "serviceregistration/kubernetes/client",
        "builtin/logical/postgresql/cmd/postgresql",
        "plugins/database/mongodb",
        "command/agent/auth/alicloud",
        "builtin/logical/rabbitmq",
        "physical/s3",
        "plugins/database/mysql",
        "http",
        "helper/testhelpers",
        "builtin/credential/aws",
        "builtin/credential/approle",
        "plugins/database/cassandra",
        "builtin/logical/ssh",
        "command/agent/template",
    ],
    deps = [
        ":google.golang.org.protobuf",
        ":github.com.golang.protobuf",
        ":github.com.hashicorp.vault.sdk_1",
        ":github.com.hashicorp.go-hclog",
        ":github.com.hashicorp.vault.api",
        ":golang.org.x.oauth2",
        ":github.com.google.go-github",
        ":github.com.hashicorp.go-cleanhttp",
        ":github.com.mitchellh.mapstructure",
        ":github.com.influxdata.influxdb",
        ":github.com.hashicorp.go-multierror",
        ":go.opentelemetry.io.otel.sdk",
        ":github.com.shirou.gopsutil",
        ":github.com.dustin.go-humanize",
        ":github.com.mitchellh.go-wordwrap",
        ":go.opentelemetry.io.otel.trace",
        ":golang.org.x.sys",
        ":go.opentelemetry.io.otel_1",
        ":github.com.hashicorp.hcl",
        ":github.com.hashicorp.consul-template",
        ":gopkg.in.mgo.v2",
        ":github.com.hashicorp.go-uuid",
        ":github.com.hashicorp.go-memdb",
        ":github.com.patrickmn.go-cache",
        ":github.com.pquerna.otp",
        ":github.com.mitchellh.cli",
        ":github.com.hashicorp.errwrap",
        ":github.com.jefferai.isbadcipher",
        ":github.com.hashicorp.consul.api",
        ":go.uber.org.atomic",
        ":golang.org.x.net",
        ":github.com.armon.go-metrics",
        ":github.com.lib.pq",
        ":go.etcd.io.etcd",
        ":github.com.coreos.go-semver",
        ":github.com.hashicorp.golang-lru",
        ":github.com.keybase.go-crypto",
        ":github.com.hashicorp.go-kms-wrapping",
        ":github.com.hashicorp.raft",
        ":golang.org.x.crypto",
        ":github.com.mitchellh.go-testing-interface",
        ":github.com.hashicorp.go-discover",
        ":github.com.armon.go-radix",
        ":google.golang.org.grpc",
        ":gopkg.in.square.go-jose.v2",
        ":github.com.mitchellh.copystructure",
        ":github.com.oklog.run",
        ":github.com.hashicorp.go-sockaddr",
        ":github.com.mitchellh.go-homedir",
        ":github.com.aws.aws-sdk-go",
        ":github.com.gocql.gocql",
        ":github.com.natefinch.atomic",
        ":github.com.aliyun.aliyun-oss-go-sdk",
        ":github.com.joyent.triton-go",
        ":github.com.hashicorp.vault-plugin-auth-oci",
        ":github.com.hashicorp.vault-plugin-secrets-gcpkms",
        ":github.com.hashicorp.vault-plugin-auth-cf",
        ":github.com.hashicorp.vault-plugin-database-snowflake",
        ":github.com.hashicorp.vault-plugin-secrets-mongodbatlas",
        ":github.com.hashicorp.vault-plugin-secrets-openldap",
        ":github.com.hashicorp.vault-plugin-auth-kerberos",
        ":github.com.hashicorp.vault-plugin-auth-centrify",
        ":github.com.hashicorp.vault-plugin-auth-kubernetes",
        ":github.com.hashicorp.vault-plugin-auth-gcp",
        ":github.com.hashicorp.vault-plugin-secrets-kv",
        ":github.com.hashicorp.vault-plugin-secrets-gcp",
        ":github.com.hashicorp.vault-plugin-auth-alicloud",
        ":github.com.hashicorp.vault-plugin-auth-jwt",
        ":github.com.hashicorp.vault-plugin-database-elasticsearch",
        ":github.com.hashicorp.vault-plugin-secrets-alicloud",
        ":github.com.hashicorp.vault-plugin-secrets-terraform",
        ":github.com.hashicorp.vault-plugin-database-mongodbatlas",
        ":github.com.hashicorp.vault-plugin-database-couchbase",
        ":github.com.hashicorp.vault-plugin-secrets-ad",
        ":github.com.hashicorp.vault-plugin-secrets-azure",
        ":github.com.hashicorp.vault-plugin-auth-azure",
        ":github.com.cockroachdb.cockroach-go",
        ":layeh.com.radius",
        ":github.com.jefferai.jsonx",
        ":github.com.mitchellh.reflectwalk",
        ":github.com.cenkalti.backoff.v3",
        ":github.com.docker.docker",
        ":github.com.docker.go-connections",
        ":github.com.oracle.oci-go-sdk",
        ":github.com.go-sql-driver.mysql",
        ":github.com.hashicorp.nomad.api",
        ":github.com.hashicorp.go-gcp-common",
        ":google.golang.org.api",
        ":go.etcd.io.bbolt",
        ":github.com.asaskevich.govalidator",
        ":github.com.ryanuber.go-glob",
        ":github.com.fatih.structs",
        ":github.com.stretchr.testify",
        ":github.com.go-test.deep",
        ":github.com.go-errors.errors",
        ":github.com.chrismalek.oktasdk-go",
        ":github.com.okta.okta-sdk-golang.v2",
        ":github.com.hashicorp.go-retryablehttp",
        ":github.com.hashicorp.go-syslog",
        ":github.com.denisenkom.go-mssqldb",
        ":cloud.google.com.go.storage",
        ":github.com.pkg.errors",
        ":github.com.Azure.azure-storage-blob-go",
        ":github.com.SAP.go-hdb",
        ":github.com.hashicorp.raft-snapshot",
        ":github.com.hashicorp.go-raftchunking",
        ":github.com.hashicorp.raft-boltdb.v2",
        ":github.com.rboyer.safeio",
        ":github.com.hashicorp.raft-autopilot",
        ":github.com.sethvargo.go-limiter",
        ":cloud.google.com.go.spanner",
        ":github.com.google.go-metrics-stackdriver",
        ":cloud.google.com.go",
        ":github.com.duosecurity.duo_api_golang",
        ":github.com.aerospike.aerospike-client-go",
        ":github.com.armon.go-proxyproto",
        ":github.com.ncw.swift",
        ":github.com.jcmturner.gokrb5.v8",
        ":github.com.kr.pretty",
        ":github.com.fatih.color",
        ":github.com.mattn.go-colorable",
        ":github.com.kr.text",
        ":github.com.ryanuber.columnize",
        ":github.com.ghodss.yaml",
        ":github.com.coreos.go-systemd",
        ":golang.org.x.term",
        ":github.com.posener.complete",
        ":github.com.mholt.archiver",
        ":github.com.Azure.go-autorest.autorest.adal",
        ":github.com.Azure.go-autorest.autorest",
        ":github.com.samuel.go-zookeeper",
        ":github.com.prometheus.client_golang",
        ":github.com.prometheus.common",
        ":go.mongodb.org.mongo-driver",
        ":github.com.aliyun.alibaba-cloud-sdk-go",
        ":github.com.michaelklishin.rabbit-hole",
        ":github.com.elazarl.go-bindata-assetfs",
        ":github.com.NYTimes.gziphandler",
    ],
)

go_module(
    name = "github.com.google.go-metrics-stackdriver",
    module = "github.com/google/go-metrics-stackdriver",
    version = "v0.2.0",
    install = [
        ".",
        "vault",
    ],
    deps = [
        ":github.com.armon.go-metrics",
        ":google.golang.org.genproto_1",
        ":google.golang.org.genproto",
        ":cloud.google.com.go_1",
        ":cloud.google.com.go",
        ":github.com.golang.protobuf",
    ],
)

go_module(
    name = "gopkg.in.resty.v1",
    module = "gopkg.in/resty.v1",
    version = "v1.12.0",
    deps = [
        ":golang.org.x.net",
    ],
)

go_module(
    name = "github.com.hashicorp.vault-plugin-database-mongodbatlas",
    module = "github.com/hashicorp/vault-plugin-database-mongodbatlas",
    version = "v0.4.0",
    deps = [
        ":github.com.hashicorp.vault.sdk_1",
        ":github.com.Sectorbob.mlab-ns2",
        ":go.mongodb.org.atlas",
        ":github.com.mitchellh.mapstructure",
    ],
)

go_module(
    name = "github.com.jcmturner.dnsutils",
    module = "github.com/jcmturner/dnsutils",
    version = "v1.0.1",
)

go_module(
    name = "github.com.mitchellh.hashstructure",
    module = "github.com/mitchellh/hashstructure",
    version = "v1.0.0",
)

go_module(
    name = "github.com.golang.protobuf",
    module = "github.com/golang/protobuf",
    version = "v1.5.2",
    install = [
        "proto",
        "ptypes",
        "ptypes/timestamp",
        "ptypes/duration",
        "ptypes/struct",
        "ptypes/any",
    ],
    deps = [
        ":google.golang.org.protobuf",
    ],
)

go_module(
    name = "github.com.digitalocean.godo",
    module = "github.com/digitalocean/godo",
    version = "v1.7.5",
    deps = [
        ":github.com.google.go-querystring",
    ],
)

go_module(
    name = "golang.org.x.xerrors",
    module = "golang.org/x/xerrors",
    version = "v0.0.0-20200804184101-5ec99f83aff1",
    install = [
        ".",
        "internal",
    ],
)

go_module(
    name = "github.com.hashicorp.vault-plugin-auth-gcp",
    module = "github.com/hashicorp/vault-plugin-auth-gcp",
    version = "v0.10.0",
    install = [
        "plugin/cache",
        "plugin",
    ],
    deps = [
        ":github.com.hashicorp.vault.sdk_1",
        ":github.com.hashicorp.go-gcp-common",
        ":github.com.hashicorp.go-hclog",
        ":google.golang.org.api",
        ":gopkg.in.square.go-jose.v2",
        ":golang.org.x.oauth2",
        ":github.com.hashicorp.go-cleanhttp",
        ":github.com.hashicorp.go-uuid",
        ":github.com.hashicorp.errwrap",
        ":github.com.hashicorp.vault.api",
    ],
)

go_module(
    name = "github.com.hashicorp.go-sockaddr",
    module = "github.com/hashicorp/go-sockaddr",
    version = "v1.0.2",
    install = [
        ".",
        "template",
    ],
    deps = [
        ":github.com.hashicorp.errwrap",
    ],
)

go_module(
    name = "github.com.aws.aws-sdk-go-v2.feature.s3.manager",
    module = "github.com/aws/aws-sdk-go-v2/feature/s3/manager",
    version = "v1.1.2",
    deps = [
        ":github.com.aws.aws-sdk-go-v2",
        ":github.com.aws.aws-sdk-go-v2.service.s3",
        ":github.com.aws.smithy-go",
    ],
)

go_module(
    name = "github.com.hashicorp.vault-plugin-secrets-mongodbatlas",
    module = "github.com/hashicorp/vault-plugin-secrets-mongodbatlas",
    version = "v0.4.0",
    deps = [
        ":github.com.hashicorp.errwrap",
        ":github.com.Sectorbob.mlab-ns2",
        ":github.com.hashicorp.vault.sdk_1",
        ":go.mongodb.org.atlas",
        ":github.com.mitchellh.mapstructure",
    ],
)

go_module(
    name = "go.uber.org.multierr",
    module = "go.uber.org/multierr",
    version = "v1.5.0",
    deps = [
        ":go.uber.org.atomic",
    ],
)

go_module(
    name = "github.com.gorilla.websocket",
    module = "github.com/gorilla/websocket",
    version = "v1.4.2",
)

go_module(
    name = "github.com.armon.go-radix",
    module = "github.com/armon/go-radix",
    version = "v1.0.0",
)

go_module(
    name = "github.com.posener.complete",
    module = "github.com/posener/complete",
    version = "v1.2.3",
    install = [
        ".",
        "cmd/install",
        "cmd",
    ],
    deps = [
        ":github.com.hashicorp.go-multierror",
    ],
)

go_module(
    name = "sigs.k8s.io.yaml",
    module = "sigs.k8s.io/yaml",
    version = "v1.2.0",
    deps = [
        ":gopkg.in.yaml.v2",
    ],
)

go_module(
    name = "github.com.xdg.scram",
    module = "github.com/xdg/scram",
    version = "v0.0.0-20180814205039-7eeb5667e42c",
    deps = [
        ":github.com.xdg.stringprep",
        ":golang.org.x.crypto",
    ],
)

go_module(
    name = "go.mongodb.org.atlas",
    module = "go.mongodb.org/atlas",
    version = "v0.7.1",
    install = [
        "mongodbatlas",
    ],
    deps = [
        ":github.com.google.go-querystring",
        ":github.com.openlyinc.pointy",
    ],
)

go_module(
    name = "github.com.mitchellh.pointerstructure",
    module = "github.com/mitchellh/pointerstructure",
    version = "v1.0.0",
    deps = [
        ":github.com.mitchellh.mapstructure",
    ],
)

go_mod_download(
    name = "go.opentelemetry.io.otel_dl",
    module = "go.opentelemetry.io/otel",
    version = "v0.20.0",
)

go_module(
    name = "go.opentelemetry.io.otel",
    module = "go.opentelemetry.io/otel",
    download = ":go.opentelemetry.io.otel_dl",
    install = [
        ".",
        "internal/global",
        "internal/trace/noop",
        "propagation",
    ],
    deps = [
        ":go.opentelemetry.io.otel.metric",
        ":go.opentelemetry.io.otel.trace",
    ],
    exported_deps = [
        ":go.opentelemetry.io.otel_1",
    ],
)

go_module(
    name = "go.opentelemetry.io.otel_1",
    module = "go.opentelemetry.io/otel",
    download = ":go.opentelemetry.io.otel_dl",
    install = [
        "unit",
        "internal",
        "codes",
        "attribute",
        "internal/baggage",
        "semconv",
    ],
)

go_module(
    name = "go.opentelemetry.io.otel.metric",
    module = "go.opentelemetry.io/otel/metric",
    version = "v0.20.0",
    install = [
        ".",
        "number",
        "registry",
    ],
    deps = [
        ":go.opentelemetry.io.otel_1",
    ],
)

go_module(
    name = "github.com.fatih.structs",
    module = "github.com/fatih/structs",
    version = "v1.1.0",
)

go_module(
    name = "golang.org.x.oauth2",
    module = "golang.org/x/oauth2",
    version = "v0.0.0-20200107190931-bf48bf16ab8d",
    install = [
        ".",
        "internal",
        "google",
        "clientcredentials",
        "jwt",
        "jws",
    ],
    deps = [
        ":golang.org.x.net",
        ":cloud.google.com.go_1",
    ],
)

go_module(
    name = "github.com.aws.aws-sdk-go-v2.service.s3",
    module = "github.com/aws/aws-sdk-go-v2/service/s3",
    version = "v1.5.0",
    install = [
        ".",
        "internal/arn",
        "internal/endpoints",
        "types",
        "internal/customizations",
    ],
    deps = [
        ":github.com.aws.aws-sdk-go-v2",
        ":github.com.aws.aws-sdk-go-v2.service.internal.s3shared",
        ":github.com.aws.aws-sdk-go-v2.service.internal.accept-encoding",
        ":github.com.aws.smithy-go",
        ":github.com.aws.aws-sdk-go-v2.service.internal.presigned-url",
    ],
)

go_module(
    name = "github.com.briankassouf.jose",
    module = "github.com/briankassouf/jose",
    version = "v0.9.2-0.20180619214549-d2569464773f",
    install = [
        ".",
        "jwt",
        "jws",
        "crypto",
    ],
)

go_module(
    name = "github.com.ghodss.yaml",
    module = "github.com/ghodss/yaml",
    version = "v1.0.1-0.20190212211648-25d852aebe32",
    deps = [
        ":gopkg.in.yaml.v2",
    ],
)

go_module(
    name = "github.com.coreos.go-semver",
    module = "github.com/coreos/go-semver",
    version = "v0.2.0",
    install = [
        "semver",
    ],
)

go_module(
    name = "github.com.Masterminds.goutils",
    module = "github.com/Masterminds/goutils",
    version = "v1.1.0",
)

go_module(
    name = "github.com.hashicorp.vault-plugin-auth-kerberos",
    module = "github.com/hashicorp/vault-plugin-auth-kerberos",
    version = "v0.4.0",
    deps = [
        ":github.com.hashicorp.vault.sdk_1",
        ":github.com.jcmturner.gokrb5.v8",
        ":github.com.hashicorp.errwrap",
        ":gopkg.in.jcmturner.goidentity.v3",
        ":github.com.hashicorp.go-hclog",
        ":github.com.hashicorp.vault.api",
    ],
)

go_module(
    name = "github.com.hashicorp.go-memdb",
    module = "github.com/hashicorp/go-memdb",
    version = "v1.0.2",
    deps = [
        ":github.com.hashicorp.go-immutable-radix",
    ],
)

go_module(
    name = "github.com.Azure.go-autorest.autorest.azure.auth",
    module = "github.com/Azure/go-autorest/autorest/azure/auth",
    version = "v0.5.7",
    deps = [
        ":github.com.dimchansky.utfbom",
        ":github.com.Azure.go-autorest.autorest.azure.cli",
        ":github.com.Azure.go-autorest.autorest",
        ":github.com.Azure.go-autorest.logger",
        ":github.com.Azure.go-autorest.autorest.adal",
    ],
)

go_module(
    name = "github.com.jackc.pgx",
    module = "github.com/jackc/pgx",
    version = "v3.3.0+incompatible",
    install = [
        ".",
        "chunkreader",
        "pgtype",
        "pgio",
        "pgproto3",
        "internal/sanitize",
    ],
    deps = [
        ":github.com.pkg.errors",
    ],
)

go_module(
    name = "github.com.aws.aws-sdk-go-v2.service.internal.presigned-url",
    module = "github.com/aws/aws-sdk-go-v2/service/internal/presigned-url",
    version = "v1.0.6",
    deps = [
        ":github.com.aws.aws-sdk-go-v2",
        ":github.com.aws.smithy-go",
    ],
)

go_module(
    name = "github.com.ryanuber.go-glob",
    module = "github.com/ryanuber/go-glob",
    version = "v1.0.0",
)

go_module(
    name = "github.com.hashicorp.go-syslog",
    module = "github.com/hashicorp/go-syslog",
    version = "v1.0.0",
)

go_module(
    name = "github.com.hashicorp.go-hclog",
    module = "github.com/hashicorp/go-hclog",
    version = "v0.16.1",
    deps = [
        ":github.com.fatih.color",
        ":github.com.mattn.go-isatty",
    ],
)

go_module(
    name = "github.com.hashicorp.go-immutable-radix",
    module = "github.com/hashicorp/go-immutable-radix",
    version = "v1.3.0",
    deps = [
        ":github.com.hashicorp.golang-lru",
    ],
)

go_module(
    name = "github.com.jcmturner.gokrb5.v8",
    module = "github.com/jcmturner/gokrb5/v8",
    version = "v8.0.0",
    install = [
        "config",
        "kadmin",
        "iana/chksumtype",
        "crypto/rfc8009",
        "crypto/rfc4757",
        "crypto",
        "iana/adtype",
        "iana/nametype",
        "types",
        "iana/errorcode",
        "iana",
        "iana/keyusage",
        "krberror",
        "crypto/rfc3961",
        "keytab",
        "client",
        "spnego",
        "iana/etypeID",
        "iana/flags",
        "service",
        "asn1tools",
        "gssapi",
        "crypto/common",
        "pac",
        "iana/patype",
        "iana/msgtype",
        "messages",
        "credentials",
        "iana/addrtype",
        "crypto/rfc3962",
        "iana/asnAppTag",
        "crypto/etype",
    ],
    deps = [
        ":github.com.jcmturner.dnsutils",
        ":github.com.jcmturner.gofork",
        ":golang.org.x.crypto",
        ":github.com.jcmturner.aescts",
        ":github.com.jcmturner.goidentity.v6",
        ":github.com.jcmturner.rpc.v2",
        ":github.com.hashicorp.go-uuid",
    ],
)

go_module(
    name = "github.com.vmware.govmomi",
    module = "github.com/vmware/govmomi",
    version = "v0.18.0",
    install = [
        ".",
        "vim25/debug",
        "list",
        "task",
        "vim25/xml",
        "vim25/progress",
        "vim25/soap",
        "session",
        "nfc",
        "vim25/methods",
        "find",
        "vim25/mo",
        "vim25/types",
        "vim25",
        "object",
        "property",
    ],
)

go_module(
    name = "github.com.hashicorp.consul-template",
    module = "github.com/hashicorp/consul-template",
    version = "v0.26.0",
    install = [
        "template",
        "signals",
        "version",
        "config",
        "manager",
        "dependency",
        "logging",
        "renderer",
        "child",
        "watch",
    ],
    deps = [
        ":github.com.BurntSushi.toml",
        ":github.com.hashicorp.consul.api",
        ":github.com.hashicorp.go-sockaddr",
        ":gopkg.in.yaml.v2",
        ":github.com.pkg.errors",
        ":github.com.davecgh.go-spew",
        ":github.com.mitchellh.mapstructure",
        ":github.com.mitchellh.go-homedir",
        ":github.com.hashicorp.hcl",
        ":github.com.hashicorp.vault.api",
        ":github.com.hashicorp.go-multierror",
        ":github.com.mattn.go-shellwords",
        ":github.com.mitchellh.hashstructure",
        ":github.com.hashicorp.go-rootcerts",
        ":github.com.hashicorp.go-syslog",
        ":github.com.hashicorp.logutils",
    ],
)

go_module(
    name = "github.com.Azure.go-autorest.autorest.to",
    module = "github.com/Azure/go-autorest/autorest/to",
    version = "v0.4.0",
)

go_module(
    name = "github.com.golang-sql.civil",
    module = "github.com/golang-sql/civil",
    version = "v0.0.0-20190719163853-cb61b32ac6fe",
)

go_module(
    name = "github.com.gorhill.cronexpr",
    module = "github.com/gorhill/cronexpr",
    version = "v0.0.0-20180427100037-88b0669f7d75",
)

go_module(
    name = "github.com.mitchellh.copystructure",
    module = "github.com/mitchellh/copystructure",
    version = "v1.0.0",
    deps = [
        ":github.com.mitchellh.reflectwalk",
    ],
)

go_module(
    name = "github.com.pquerna.otp",
    module = "github.com/pquerna/otp",
    version = "v1.2.1-0.20191009055518-468c2dd2b58d",
    install = [
        ".",
        "totp",
        "hotp",
    ],
    deps = [
        ":github.com.boombuler.barcode",
    ],
)

go_module(
    name = "go.mongodb.org.mongo-driver",
    module = "go.mongodb.org/mongo-driver",
    version = "v1.4.6",
    install = [
        "bson/bsonoptions",
        "tag",
        "x/mongo/driver/description",
        "event",
        "bson",
        "mongo/options",
        "x/mongo/driver/mongocrypt",
        "x/mongo/driver/dns",
        "x/bsonx/bsoncore",
        "x/mongo/driver/topology",
        "x/mongo/driver/ocsp",
        "mongo/writeconcern",
        "mongo",
        "x/mongo/driver/connstring",
        "x/mongo/driver/mongocrypt/options",
        "x/mongo/driver/address",
        "x/mongo/driver/session",
        "bson/bsoncodec",
        "mongo/readconcern",
        "x/mongo/driver",
        "internal",
        "mongo/readpref",
        "x/mongo/driver/auth",
        "x/mongo/driver/wiremessage",
        "bson/bsontype",
        "version",
        "bson/bsonrw",
        "bson/primitive",
        "x/mongo/driver/operation",
        "x/mongo/driver/uuid",
        "x/bsonx",
    ],
    deps = [
        ":github.com.go-stack.stack",
        ":golang.org.x.sync",
        ":golang.org.x.crypto",
        ":github.com.pkg.errors",
        ":github.com.klauspost.compress",
        ":github.com.golang.snappy",
        ":github.com.xdg.scram",
        ":github.com.xdg.stringprep",
        ":github.com.aws.aws-sdk-go",
    ],
)

go_module(
    name = "github.com.hashicorp.go-secure-stdlib.strutil",
    module = "github.com/hashicorp/go-secure-stdlib/strutil",
    version = "v0.1.1",
    deps = [
        ":github.com.ryanuber.go-glob",
    ],
)

go_module(
    name = "github.com.jmespath.go-jmespath",
    module = "github.com/jmespath/go-jmespath",
    version = "v0.4.0",
)

go_module(
    name = "go.etcd.io.bbolt",
    module = "go.etcd.io/bbolt",
    version = "v1.3.5",
)

go_module(
    name = "github.com.hashicorp.cap",
    module = "github.com/hashicorp/cap",
    version = "v0.1.0",
    install = [
        "oidc",
        "jwt",
        "oidc/internal/base62",
        "oidc/internal/strutils",
    ],
    deps = [
        ":gopkg.in.square.go-jose.v2",
        ":github.com.coreos.go-oidc.v3",
        ":github.com.stretchr.testify",
        ":golang.org.x.oauth2",
        ":github.com.hashicorp.go-hclog",
        ":golang.org.x.text",
        ":github.com.hashicorp.go-cleanhttp",
        ":github.com.hashicorp.go-uuid",
    ],
)

go_module(
    name = "github.com.go-yaml.yaml",
    module = "github.com/go-yaml/yaml",
    version = "v2.1.0+incompatible",
)

go_module(
    name = "github.com.ncw.swift",
    module = "github.com/ncw/swift",
    version = "v1.0.47",
)

go_module(
    name = "github.com.hashicorp.logutils",
    module = "github.com/hashicorp/logutils",
    version = "v1.0.0",
)

go_mod_download(
    name = "github.com.hashicorp.vault.sdk_dl",
    module = "github.com/hashicorp/vault/sdk",
    version = "v0.2.2-0.20210825150427-9b1f4d486f5d",
)

go_module(
    name = "github.com.hashicorp.vault.sdk",
    module = "github.com/hashicorp/vault/sdk",
    download = ":github.com.hashicorp.vault.sdk_dl",
    install = [
        "testing/stepwise",
        "testing/stepwise/environments/docker",
    ],
    deps = [
        ":github.com.hashicorp.go-hclog",
        ":github.com.hashicorp.errwrap",
        ":github.com.hashicorp.vault.api",
        ":github.com.hashicorp.go-uuid",
        ":github.com.docker.docker",
        ":github.com.hashicorp.go-cleanhttp",
        ":github.com.docker.go-connections",
        ":github.com.hashicorp.go-multierror",
        ":golang.org.x.net",
    ],
    exported_deps = [
        ":github.com.hashicorp.vault.sdk_1",
    ],
)

go_module(
    name = "github.com.hashicorp.vault.sdk_1",
    module = "github.com/hashicorp/vault/sdk",
    download = ":github.com.hashicorp.vault.sdk_dl",
    install = [
        "logical",
        "helper/cidrutil",
        "database/dbplugin/v5",
        "helper/compressutil",
        "helper/dbtxn",
        "database/dbplugin/v5/proto",
        "queue",
        "physical/file",
        "framework",
        "helper/consts",
        "helper/parseutil",
        "physical",
        "helper/locksutil",
        "database/helper/connutil",
        "version",
        "helper/template",
        "helper/kdf",
        "helper/hclutil",
        "plugin/pb",
        "helper/mlock",
        "database/dbplugin",
        "database/helper/credsutil",
        "helper/keysutil",
        "helper/useragent",
        "helper/jsonutil",
        "helper/wrapping",
        "helper/logging",
        "helper/pluginutil",
        "database/dbplugin/v5/testing",
        "helper/pathmanager",
        "plugin/mock",
        "helper/certutil",
        "helper/policyutil",
        "helper/errutil",
        "helper/ldaputil",
        "helper/authmetadata",
        "helper/awsutil",
        "helper/cryptoutil",
        "helper/tokenutil",
        "helper/base62",
        "plugin",
        "helper/salt",
        "helper/identitytpl",
        "physical/inmem",
        "helper/strutil",
        "helper/tlsutil",
        "helper/license",
        "database/helper/dbutil",
        "helper/password",
        "helper/pointerutil",
    ],
    deps = [
        ":github.com.mitchellh.go-testing-interface",
        ":github.com.hashicorp.errwrap",
        ":github.com.hashicorp.go-multierror",
        ":google.golang.org.protobuf",
        ":github.com.hashicorp.go-sockaddr",
        ":github.com.golang.protobuf",
        ":github.com.mitchellh.copystructure",
        ":github.com.hashicorp.go-hclog",
        ":github.com.hashicorp.go-plugin",
        ":github.com.armon.go-metrics",
        ":google.golang.org.grpc",
        ":github.com.pierrec.lz4",
        ":github.com.golang.snappy",
        ":github.com.ryanuber.go-glob",
        ":github.com.mitchellh.mapstructure",
        ":github.com.hashicorp.go-uuid",
        ":github.com.hashicorp.go-kms-wrapping.entropy",
        ":github.com.hashicorp.golang-lru",
        ":go.uber.org.atomic",
        ":github.com.hashicorp.hcl",
        ":golang.org.x.sys",
        ":golang.org.x.crypto",
        ":github.com.hashicorp.go-version",
        ":github.com.hashicorp.go-immutable-radix",
        ":github.com.go-ldap.ldap.v3",
        ":github.com.aws.aws-sdk-go",
        ":github.com.pkg.errors",
        ":github.com.armon.go-radix",
    ],
)

go_module(
    name = "github.com.denisenkom.go-mssqldb",
    module = "github.com/denisenkom/go-mssqldb",
    version = "v0.0.0-20200428022330-06a60b6afbbc",
    install = [
        ".",
        "internal/cp",
        "internal/querytext",
        "internal/decimal",
    ],
    deps = [
        ":github.com.golang-sql.civil",
        ":golang.org.x.crypto",
    ],
)

go_module(
    name = "github.com.go-ldap.ldif",
    module = "github.com/go-ldap/ldif",
    version = "v0.0.0-20200320164324-fd88d9b715b3",
    deps = [
        ":github.com.go-ldap.ldap.v3",
    ],
)

go_module(
    name = "github.com.hashicorp.vault-plugin-secrets-gcp",
    module = "github.com/hashicorp/vault-plugin-secrets-gcp",
    version = "v0.10.2",
    install = [
        "plugin/iamutil",
        "plugin",
        "plugin/util",
    ],
    deps = [
        ":github.com.hashicorp.go-gcp-common",
        ":github.com.hashicorp.errwrap",
        ":google.golang.org.api",
        ":github.com.hashicorp.go-multierror",
        ":github.com.mitchellh.mapstructure",
        ":golang.org.x.oauth2",
        ":github.com.hashicorp.go-cleanhttp",
        ":github.com.hashicorp.vault.sdk_1",
        ":github.com.hashicorp.vault-plugin-auth-gcp",
        ":github.com.hashicorp.hcl",
    ],
)

go_module(
    name = "github.com.hashicorp.vault-plugin-database-elasticsearch",
    module = "github.com/hashicorp/vault-plugin-database-elasticsearch",
    version = "v0.8.0",
    deps = [
        ":github.com.hashicorp.errwrap",
        ":github.com.hashicorp.go-multierror",
        ":github.com.hashicorp.go-retryablehttp",
        ":github.com.hashicorp.vault.sdk_1",
        ":github.com.hashicorp.go-rootcerts",
    ],
)

go_module(
    name = "github.com.centrify.cloud-golang-sdk",
    module = "github.com/centrify/cloud-golang-sdk",
    version = "v0.0.0-20190214225812-119110094d0f",
    install = [
        "restapi",
        "oauth",
    ],
)

go_module(
    name = "github.com.modern-go.reflect2",
    module = "github.com/modern-go/reflect2",
    version = "v1.0.1",
    deps = [
        ":github.com.modern-go.concurrent",
    ],
)

go_module(
    name = "github.com.hashicorp.go-gcp-common",
    module = "github.com/hashicorp/go-gcp-common",
    version = "v0.7.0",
    install = [
        "gcputil",
    ],
    deps = [
        ":google.golang.org.api",
        ":golang.org.x.oauth2",
        ":github.com.mitchellh.go-homedir",
        ":github.com.hashicorp.go-cleanhttp",
    ],
)

go_module(
    name = "github.com.stretchr.testify",
    module = "github.com/stretchr/testify",
    version = "v1.7.0",
    install = [
        "require",
        "assert",
    ],
    deps = [
        ":github.com.pmezard.go-difflib",
        ":gopkg.in.yaml.v3",
        ":github.com.davecgh.go-spew",
    ],
)

go_module(
    name = "k8s.io.apimachinery",
    module = "k8s.io/apimachinery",
    version = "v0.22.1",
    install = [
        "pkg/runtime/serializer/recognizer",
        "pkg/util/runtime",
        "pkg/apis/meta/internalversion/scheme",
        "pkg/version",
        "pkg/fields",
        "pkg/util/managedfields",
        "pkg/util/remotecommand",
        "pkg/util/httpstream/spdy",
        "pkg/apis/meta/v1beta1",
        "pkg/runtime/serializer/streaming",
        "pkg/runtime/schema",
        "pkg/runtime/serializer/versioning",
        "pkg/util/framer",
        "pkg/apis/meta/v1/unstructured",
        "pkg/util/yaml",
        "pkg/util/naming",
        "pkg/labels",
        "pkg/types",
        "third_party/forked/golang/json",
        "third_party/forked/golang/reflect",
        "pkg/util/sets",
        "pkg/util/diff",
        "pkg/util/wait",
        "pkg/runtime/serializer",
        "pkg/runtime/serializer/protobuf",
        "pkg/util/cache",
        "pkg/util/intstr",
        "pkg/util/validation",
        "pkg/util/clock",
        "pkg/api/meta",
        "pkg/util/strategicpatch",
        "pkg/selection",
        "pkg/conversion",
        "pkg/util/httpstream",
        "pkg/api/resource",
        "pkg/util/errors",
        "pkg/conversion/queryparams",
        "pkg/util/mergepatch",
        "pkg/watch",
        "pkg/api/errors",
        "pkg/apis/meta/internalversion",
        "pkg/util/json",
        "pkg/util/validation/field",
        "pkg/apis/meta/v1",
        "pkg/runtime",
        "third_party/forked/golang/netutil",
        "pkg/util/net",
        "pkg/runtime/serializer/json",
    ],
    deps = [
        ":k8s.io.klog.v2",
        ":sigs.k8s.io.structured-merge-diff.v4",
        ":github.com.moby.spdystream",
        ":github.com.gogo.protobuf",
        ":sigs.k8s.io.yaml",
        ":github.com.google.go-cmp",
        ":github.com.davecgh.go-spew",
        ":github.com.google.gofuzz",
        ":k8s.io.kube-openapi",
        ":gopkg.in.inf.v0",
        ":golang.org.x.net",
        ":github.com.modern-go.reflect2",
        ":github.com.json-iterator.go",
    ],
)

go_module(
    name = "github.com.ryanuber.columnize",
    module = "github.com/ryanuber/columnize",
    version = "v2.1.0+incompatible",
)

go_module(
    name = "github.com.hashicorp.hcl",
    module = "github.com/hashicorp/hcl",
    version = "v1.0.1-vault-3",
    install = [
        ".",
        "hcl/strconv",
        "json/token",
        "hcl/ast",
        "hcl/token",
        "json/parser",
        "hcl/scanner",
        "hcl/parser",
        "hcl/printer",
        "json/scanner",
    ],
)

go_module(
    name = "github.com.okta.okta-sdk-golang",
    module = "github.com/okta/okta-sdk-golang",
    version = "v1.1.0",
    install = [
        "okta/cache",
    ],
    deps = [
        ":github.com.patrickmn.go-cache",
    ],
)

go_module(
    name = "github.com.sirupsen.logrus",
    module = "github.com/sirupsen/logrus",
    version = "v1.7.0",
    deps = [
        ":golang.org.x.sys",
    ],
)

go_module(
    name = "github.com.coreos.go-oidc.v3",
    module = "github.com/coreos/go-oidc/v3",
    version = "v3.0.0",
    install = [
        "oidc",
    ],
    deps = [
        ":gopkg.in.square.go-jose.v2",
        ":golang.org.x.oauth2",
    ],
)

go_module(
    name = "github.com.aliyun.aliyun-oss-go-sdk",
    module = "github.com/aliyun/aliyun-oss-go-sdk",
    version = "v0.0.0-20190307165228-86c17b95fcd5",
    install = [
        "oss",
    ],
    deps = [
        ":golang.org.x.time",
    ],
)

go_module(
    name = "github.com.hashicorp.errwrap",
    module = "github.com/hashicorp/errwrap",
    version = "v1.1.0",
)

go_module(
    name = "github.com.pkg.errors",
    module = "github.com/pkg/errors",
    version = "v0.9.1",
)

go_module(
    name = "github.com.michaelklishin.rabbit-hole",
    module = "github.com/michaelklishin/rabbit-hole",
    version = "v0.0.0-20191008194146-93d9988f0cd5",
)

go_module(
    name = "github.com.pierrec.lz4",
    module = "github.com/pierrec/lz4",
    version = "v2.5.2+incompatible",
    install = [
        ".",
        "internal/xxh32",
    ],
)

go_module(
    name = "github.com.Azure.azure-storage-blob-go",
    module = "github.com/Azure/azure-storage-blob-go",
    version = "v0.13.0",
    install = [
        "azblob",
    ],
    deps = [
        ":github.com.google.uuid",
        ":github.com.Azure.azure-pipeline-go",
    ],
)

go_module(
    name = "github.com.opencontainers.image-spec",
    module = "github.com/opencontainers/image-spec",
    version = "v1.0.1",
    install = [
        "specs-go/v1",
        "specs-go",
    ],
    deps = [
        ":github.com.opencontainers.go-digest",
    ],
)

go_mod_download(
    name = "cloud.google.com.go_dl",
    module = "cloud.google.com/go",
    version = "v0.56.0",
)

go_module(
    name = "cloud.google.com.go",
    module = "cloud.google.com/go",
    download = ":cloud.google.com.go_dl",
    install = [
        "internal",
        "internal/trace",
        "iam",
        "monitoring/apiv3",
        "kms/apiv1",
        "longrunning/autogen",
        "longrunning",
    ],
    deps = [
        ":google.golang.org.api",
        ":google.golang.org.grpc",
        ":github.com.googleapis.gax-go.v2",
        ":go.opencensus.io",
        ":google.golang.org.genproto_1",
        ":google.golang.org.genproto",
        ":github.com.golang.protobuf",
    ],
    exported_deps = [
        ":cloud.google.com.go_1",
    ],
)

go_module(
    name = "cloud.google.com.go_1",
    module = "cloud.google.com/go",
    download = ":cloud.google.com.go_dl",
    install = [
        "internal/optional",
        "civil",
        "internal/protostruct",
        "compute/metadata",
        "internal/fields",
        "internal/version",
    ],
    deps = [
        ":github.com.golang.protobuf",
    ],
)

go_module(
    name = "github.com.SAP.go-hdb",
    module = "github.com/SAP/go-hdb",
    version = "v0.14.1",
    install = [
        "internal/bufio",
        "driver",
        "driver/sqltrace",
        "internal/protocol",
        "internal/unicode/cesu8",
        "internal/unicode",
    ],
    deps = [
        ":golang.org.x.text",
    ],
)

go_module(
    name = "github.com.mitchellh.go-homedir",
    module = "github.com/mitchellh/go-homedir",
    version = "v1.1.0",
)

go_module(
    name = "github.com.gocql.gocql",
    module = "github.com/gocql/gocql",
    version = "v0.0.0-20210401103645-80ab1e13e309",
    install = [
        ".",
        "internal/lru",
        "internal/murmur",
        "internal/streams",
    ],
    deps = [
        ":github.com.hailocab.go-hostpool",
        ":gopkg.in.inf.v0",
        ":github.com.golang.snappy",
    ],
)

go_module(
    name = "cloud.google.com.go.storage",
    module = "cloud.google.com/go/storage",
    version = "v1.6.0",
    deps = [
        ":google.golang.org.api",
        ":cloud.google.com.go_1",
        ":cloud.google.com.go",
        ":google.golang.org.genproto",
        ":google.golang.org.genproto_1",
        ":github.com.googleapis.gax-go.v2",
    ],
)

go_module(
    name = "gopkg.in.ini.v1",
    module = "gopkg.in/ini.v1",
    version = "v1.42.0",
)

go_module(
    name = "github.com.Azure.go-ntlmssp",
    module = "github.com/Azure/go-ntlmssp",
    version = "v0.0.0-20200615164410-66371956d46c",
    deps = [
        ":golang.org.x.crypto",
    ],
)

go_module(
    name = "k8s.io.api",
    module = "k8s.io/api",
    version = "v0.22.1",
    install = [
        "autoscaling/v1",
        "authorization/v1beta1",
        "core/v1",
        "networking/v1",
        "autoscaling/v2beta1",
        "apps/v1",
        "authentication/v1beta1",
        "apiserverinternal/v1alpha1",
        "rbac/v1beta1",
        "storage/v1alpha1",
        "events/v1beta1",
        "authentication/v1",
        "networking/v1beta1",
        "discovery/v1",
        "coordination/v1",
        "autoscaling/v2beta2",
        "certificates/v1",
        "rbac/v1alpha1",
        "scheduling/v1",
        "policy/v1beta1",
        "node/v1beta1",
        "flowcontrol/v1alpha1",
        "apps/v1beta2",
        "admissionregistration/v1",
        "policy/v1",
        "coordination/v1beta1",
        "storage/v1beta1",
        "rbac/v1",
        "node/v1",
        "scheduling/v1alpha1",
        "discovery/v1beta1",
        "imagepolicy/v1alpha1",
        "storage/v1",
        "extensions/v1beta1",
        "apps/v1beta1",
        "batch/v1beta1",
        "admissionregistration/v1beta1",
        "certificates/v1beta1",
        "batch/v1",
        "scheduling/v1beta1",
        "authorization/v1",
        "flowcontrol/v1beta1",
        "events/v1",
        "node/v1alpha1",
    ],
    deps = [
        ":k8s.io.apimachinery",
        ":github.com.gogo.protobuf",
    ],
)

go_module(
    name = "go.etcd.io.etcd",
    module = "go.etcd.io/etcd",
    version = "v0.5.0-alpha.5.0.20200425165423-262c93980547",
    install = [
        "client",
        "raft/confchange",
        "version",
        "pkg/types",
        "raft/raftpb",
        "mvcc/mvccpb",
        "etcdserver/etcdserverpb",
        "clientv3/balancer/connectivity",
        "pkg/systemd",
        "clientv3/balancer",
        "raft",
        "etcdserver/api/membership/membershippb",
        "raft/tracker",
        "pkg/pathutil",
        "pkg/logutil",
        "raft/quorum",
        "clientv3/credentials",
        "etcdserver/api/v3rpc/rpctypes",
        "pkg/srv",
        "pkg/fileutil",
        "clientv3/balancer/picker",
        "pkg/transport",
        "clientv3/concurrency",
        "clientv3",
        "pkg/tlsutil",
        "clientv3/balancer/resolver/endpoint",
        "auth/authpb",
    ],
    deps = [
        ":github.com.modern-go.reflect2",
        ":github.com.json-iterator.go",
        ":github.com.coreos.go-semver",
        ":github.com.golang.protobuf",
        ":github.com.gogo.protobuf",
        ":google.golang.org.grpc",
        ":golang.org.x.net",
        ":go.uber.org.zap",
        ":github.com.coreos.go-systemd.v22",
        ":github.com.google.uuid",
    ],
)

go_module(
    name = "github.com.peterbourgon.diskv",
    module = "github.com/peterbourgon/diskv",
    version = "v2.0.1+incompatible",
    deps = [
        ":github.com.google.btree",
    ],
)

go_module(
    name = "gopkg.in.yaml.v2",
    module = "gopkg.in/yaml.v2",
    version = "v2.4.0",
)

go_module(
    name = "github.com.Azure.go-autorest.tracing",
    module = "github.com/Azure/go-autorest/tracing",
    version = "v0.6.0",
)

go_module(
    name = "github.com.Jeffail.gabs",
    module = "github.com/Jeffail/gabs",
    version = "v1.1.1",
)

go_module(
    name = "gopkg.in.jcmturner.goidentity.v3",
    module = "gopkg.in/jcmturner/goidentity.v3",
    version = "v3.0.0",
    deps = [
        ":github.com.hashicorp.go-uuid",
    ],
)

go_module(
    name = "k8s.io.kube-openapi",
    module = "k8s.io/kube-openapi",
    version = "v0.0.0-20210421082810-95288971da7e",
    install = [
        "pkg/util/proto",
    ],
    deps = [
        ":github.com.googleapis.gnostic",
        ":gopkg.in.yaml.v2",
    ],
)

go_module(
    name = "github.com.Masterminds.semver",
    module = "github.com/Masterminds/semver",
    version = "v1.5.0",
)

go_module(
    name = "github.com.hailocab.go-hostpool",
    module = "github.com/hailocab/go-hostpool",
    version = "v0.0.0-20160125115350-e80d13ce29ed",
)

go_module(
    name = "github.com.kelseyhightower.envconfig",
    module = "github.com/kelseyhightower/envconfig",
    version = "v1.4.0",
)

go_module(
    name = "github.com.hashicorp.go-kms-wrapping",
    module = "github.com/hashicorp/go-kms-wrapping",
    version = "v0.5.16",
    install = [
        ".",
        "wrappers/gcpckms",
        "internal/xor",
        "wrappers/alicloudkms",
        "wrappers/aead",
        "wrappers/transit",
        "wrappers/awskms",
        "wrappers/ocikms",
        "wrappers/azurekeyvault",
    ],
    deps = [
        ":golang.org.x.net",
        ":google.golang.org.api",
        ":cloud.google.com.go",
        ":google.golang.org.genproto",
        ":github.com.aliyun.alibaba-cloud-sdk-go",
        ":github.com.golang.protobuf",
        ":google.golang.org.protobuf",
        ":github.com.hashicorp.go-hclog",
        ":github.com.hashicorp.go-uuid",
        ":golang.org.x.crypto",
        ":github.com.hashicorp.vault.api",
        ":github.com.aws.aws-sdk-go",
        ":github.com.hashicorp.go-cleanhttp",
        ":github.com.hashicorp.vault.sdk_1",
        ":github.com.oracle.oci-go-sdk",
        ":github.com.Azure.go-autorest.autorest.to",
        ":github.com.Azure.go-autorest.autorest.azure.auth",
        ":github.com.Azure.azure-sdk-for-go",
        ":github.com.Azure.go-autorest.autorest",
    ],
)

go_module(
    name = "github.com.hashicorp.raft-boltdb.v2",
    module = "github.com/hashicorp/raft-boltdb/v2",
    version = "v2.0.0-20210421194847-a7e34179d62c",
    deps = [
        ":github.com.hashicorp.go-msgpack",
        ":github.com.hashicorp.raft",
        ":go.etcd.io.bbolt",
    ],
)

go_module(
    name = "github.com.cockroachdb.cockroach-go",
    module = "github.com/cockroachdb/cockroach-go",
    version = "v0.0.0-20181001143604-e0a95dfd547c",
    install = [
        "crdb",
    ],
    deps = [
        ":github.com.jackc.pgx",
        ":github.com.lib.pq",
    ],
)

go_module(
    name = "github.com.gogo.protobuf",
    module = "github.com/gogo/protobuf",
    version = "v1.3.2",
    install = [
        "proto",
        "sortkeys",
        "gogoproto",
        "protoc-gen-gogo/descriptor",
    ],
)

go_module(
    name = "github.com.natefinch.atomic",
    module = "github.com/natefinch/atomic",
    version = "v0.0.0-20150920032501-a62ce929ffcc",
)

go_module(
    name = "github.com.tklauser.numcpus",
    module = "github.com/tklauser/numcpus",
    version = "v0.2.2",
)

go_module(
    name = "github.com.mitchellh.mapstructure",
    module = "github.com/mitchellh/mapstructure",
    version = "v1.3.3",
)

go_module(
    name = "github.com.svanharmelen.jsonapi",
    module = "github.com/svanharmelen/jsonapi",
    version = "v0.0.0-20180618144545-0c0828c3f16d",
)

go_module(
    name = "golang.org.x.sys",
    module = "golang.org/x/sys",
    version = "v0.0.0-20210616094352-59db8d763f22",
    install = [
        "unix",
        "cpu",
        "internal/unsafeheader",
    ],
)

go_module(
    name = "github.com.aws.aws-sdk-go-v2",
    module = "github.com/aws/aws-sdk-go-v2",
    version = "v1.3.2",
    install = [
        "internal/endpoints",
        "aws/signer/internal/v4",
        "internal/timeconv",
        "aws/middleware",
        "aws/transport/http",
        "aws/arn",
        "aws",
        "internal/sdkio",
        "internal/rand",
        "internal/strings",
        "aws/signer/v4",
        "internal/awsutil",
        "aws/ratelimit",
        "internal/sdk",
        "aws/retry",
        "internal/sync/singleflight",
    ],
    deps = [
        ":github.com.aws.smithy-go",
        ":github.com.jmespath.go-jmespath",
    ],
)

go_module(
    name = "github.com.cenkalti.backoff.v3",
    module = "github.com/cenkalti/backoff/v3",
    version = "v3.0.0",
)

go_module(
    name = "go.opentelemetry.io.otel.sdk",
    module = "go.opentelemetry.io/otel/sdk",
    version = "v0.20.0",
    install = [
        "instrumentation",
        "trace",
        "internal",
        "resource",
    ],
    deps = [
        ":go.opentelemetry.io.otel_1",
        ":go.opentelemetry.io.otel.trace",
        ":go.opentelemetry.io.otel",
    ],
)

go_module(
    name = "go.opentelemetry.io.otel.trace",
    module = "go.opentelemetry.io/otel/trace",
    version = "v0.20.0",
    deps = [
        ":go.opentelemetry.io.otel_1",
    ],
)

go_module(
    name = "github.com.dsnet.compress",
    module = "github.com/dsnet/compress",
    version = "v0.0.1",
    install = [
        ".",
        "internal/prefix",
        "bzip2",
        "internal/errors",
        "bzip2/internal/sais",
        "internal",
    ],
)

go_module(
    name = "github.com.go-errors.errors",
    module = "github.com/go-errors/errors",
    version = "v1.0.1",
)

go_module(
    name = "github.com.asaskevich.govalidator",
    module = "github.com/asaskevich/govalidator",
    version = "v0.0.0-20190424111038-f61b66f89f4a",
)

go_module(
    name = "sigs.k8s.io.structured-merge-diff.v4",
    module = "sigs.k8s.io/structured-merge-diff/v4",
    version = "v4.1.2",
    install = [
        "value",
        "fieldpath",
        "typed",
        "schema",
    ],
    deps = [
        ":gopkg.in.yaml.v2",
        ":github.com.json-iterator.go",
    ],
)

go_module(
    name = "github.com.hashicorp.mdns",
    module = "github.com/hashicorp/mdns",
    version = "v1.0.1",
    deps = [
        ":github.com.miekg.dns",
        ":golang.org.x.net",
    ],
)

go_module(
    name = "github.com.klauspost.compress",
    module = "github.com/klauspost/compress",
    version = "v1.9.5",
    install = [
        "zstd",
        "huff0",
        "fse",
        "snappy",
        "zstd/internal/xxhash",
    ],
)

go_module(
    name = "github.com.evanphx.json-patch",
    module = "github.com/evanphx/json-patch",
    version = "v4.11.0+incompatible",
    deps = [
        ":github.com.pkg.errors",
    ],
)

go_module(
    name = "github.com.google.go-querystring",
    module = "github.com/google/go-querystring",
    version = "v1.0.0",
    install = [
        "query",
    ],
)

go_module(
    name = "github.com.hashicorp.go-rootcerts",
    module = "github.com/hashicorp/go-rootcerts",
    version = "v1.0.2",
)

go_module(
    name = "github.com.jcmturner.aescts",
    module = "github.com/jcmturner/aescts",
    version = "v1.0.1",
)

go_module(
    name = "github.com.fatih.color",
    module = "github.com/fatih/color",
    version = "v1.11.0",
    deps = [
        ":github.com.mattn.go-colorable",
        ":github.com.mattn.go-isatty",
    ],
)

go_module(
    name = "github.com.mattn.go-isatty",
    module = "github.com/mattn/go-isatty",
    version = "v0.0.12",
    deps = [
        ":golang.org.x.sys",
    ],
)

go_module(
    name = "github.com.keybase.go-crypto",
    module = "github.com/keybase/go-crypto",
    version = "v0.0.0-20190403132359-d65b6b94177f",
    install = [
        "openpgp",
        "ed25519/internal/edwards25519",
        "openpgp/ecdh",
        "curve25519",
        "brainpool",
        "openpgp/errors",
        "openpgp/packet",
        "rsa",
        "openpgp/s2k",
        "cast5",
        "openpgp/armor",
        "openpgp/elgamal",
        "ed25519",
    ],
)

go_module(
    name = "github.com.hashicorp.go-discover",
    module = "github.com/hashicorp/go-discover",
    version = "v0.0.0-20201029210230-738cb3105cd0",
    install = [
        ".",
        "provider/triton",
        "provider/linode",
        "provider/os",
        "provider/tencentcloud",
        "provider/packet",
        "provider/k8s",
        "provider/digitalocean",
        "provider/gce",
        "provider/aliyun",
        "provider/scaleway",
        "provider/mdns",
        "provider/aws",
        "provider/azure",
        "provider/softlayer",
        "provider/vsphere",
    ],
    deps = [
        ":github.com.joyent.triton-go",
        ":github.com.linode.linodego",
        ":golang.org.x.oauth2",
        ":github.com.gophercloud.gophercloud",
        ":github.com.tencentcloud.tencentcloud-sdk-go",
        ":github.com.packethost.packngo",
        ":k8s.io.api",
        ":k8s.io.client-go",
        ":github.com.mitchellh.go-homedir",
        ":k8s.io.apimachinery",
        ":github.com.hashicorp.go-multierror",
        ":github.com.digitalocean.godo",
        ":google.golang.org.api",
        ":github.com.denverdino.aliyungo",
        ":github.com.nicolai86.scaleway-sdk",
        ":github.com.hashicorp.mdns",
        ":github.com.aws.aws-sdk-go",
        ":github.com.Azure.go-autorest.autorest.azure.auth",
        ":github.com.Azure.azure-sdk-for-go",
        ":github.com.Azure.go-autorest.autorest",
        ":github.com.softlayer.softlayer-go",
        ":github.com.vmware.govmomi",
        ":github.com.hashicorp.vic",
    ],
)

go_module(
    name = "github.com.mholt.archiver",
    module = "github.com/mholt/archiver",
    version = "v3.1.1+incompatible",
    deps = [
        ":github.com.golang.snappy",
        ":github.com.dsnet.compress",
        ":github.com.pierrec.lz4",
        ":github.com.nwaples.rardecode",
        ":github.com.xi2.xz",
        ":github.com.ulikunitz.xz",
    ],
)

go_module(
    name = "github.com.xdg.stringprep",
    module = "github.com/xdg/stringprep",
    version = "v1.0.0",
    deps = [
        ":golang.org.x.text",
    ],
)

go_module(
    name = "github.com.aws.aws-sdk-go-v2.credentials",
    module = "github.com/aws/aws-sdk-go-v2/credentials",
    version = "v1.1.5",
    deps = [
        ":github.com.aws.aws-sdk-go-v2",
    ],
)

go_module(
    name = "github.com.hashicorp.vault-plugin-auth-jwt",
    module = "github.com/hashicorp/vault-plugin-auth-jwt",
    version = "v0.10.1",
    deps = [
        ":github.com.hashicorp.vault.sdk_1",
        ":github.com.hashicorp.cap",
        ":github.com.ryanuber.go-glob",
        ":github.com.hashicorp.go-hclog",
        ":golang.org.x.oauth2",
        ":github.com.hashicorp.go-cleanhttp",
        ":google.golang.org.api",
        ":github.com.hashicorp.errwrap",
        ":gopkg.in.square.go-jose.v2",
        ":github.com.patrickmn.go-cache",
        ":github.com.hashicorp.go-sockaddr",
        ":github.com.mitchellh.pointerstructure",
        ":github.com.hashicorp.vault.api",
        ":github.com.mitchellh.mapstructure",
    ],
)

go_module(
    name = "github.com.hashicorp.yamux",
    module = "github.com/hashicorp/yamux",
    version = "v0.0.0-20181012175058-2f1d1f20f75d",
)

go_module(
    name = "github.com.openlyinc.pointy",
    module = "github.com/openlyinc/pointy",
    version = "v1.1.2",
)

go_module(
    name = "github.com.armon.go-metrics",
    module = "github.com/armon/go-metrics",
    version = "v0.3.7",
    install = [
        ".",
        "prometheus",
        "datadog",
        "circonus",
    ],
    deps = [
        ":github.com.prometheus.client_golang",
        ":github.com.hashicorp.go-immutable-radix",
        ":github.com.DataDog.datadog-go",
        ":github.com.circonus-labs.circonus-gometrics",
    ],
)

go_module(
    name = "github.com.duosecurity.duo_api_golang",
    module = "github.com/duosecurity/duo_api_golang",
    version = "v0.0.0-20190308151101-6c680f768e74",
    install = [
        ".",
        "authapi",
    ],
)

go_module(
    name = "go.uber.org.atomic",
    module = "go.uber.org/atomic",
    version = "v1.6.0",
)

go_module(
    name = "github.com.opencontainers.go-digest",
    module = "github.com/opencontainers/go-digest",
    version = "v1.0.0",
)

go_module(
    name = "github.com.go-sql-driver.mysql",
    module = "github.com/go-sql-driver/mysql",
    version = "v1.5.0",
)

go_module(
    name = "github.com.jeffchao.backoff",
    module = "github.com/jeffchao/backoff",
    version = "v0.0.0-20140404060208-9d7fd7aa17f2",
)

go_mod_download(
    name = "google.golang.org.genproto_dl",
    module = "google.golang.org/genproto",
    version = "v0.0.0-20201019141844-1ed22bb0c154",
)

go_module(
    name = "google.golang.org.genproto",
    module = "google.golang.org/genproto",
    download = ":google.golang.org.genproto_dl",
    install = [
        "googleapis/monitoring/v3",
        "googleapis/longrunning",
        "googleapis/iam/v1",
        "googleapis/spanner/admin/instance/v1",
        "googleapis/spanner/v1",
        "googleapis/cloud/kms/v1",
    ],
    deps = [
        ":google.golang.org.protobuf",
        ":github.com.golang.protobuf",
        ":google.golang.org.grpc",
    ],
    exported_deps = [
        ":google.golang.org.genproto_1",
    ],
)

go_module(
    name = "google.golang.org.genproto_1",
    module = "google.golang.org/genproto",
    download = ":google.golang.org.genproto_dl",
    install = [
        "googleapis/rpc/errdetails",
        "googleapis/api/distribution",
        "googleapis/rpc/code",
        "googleapis/api/metric",
        "googleapis/type/calendarperiod",
        "googleapis/api/label",
        "googleapis/rpc/status",
        "googleapis/type/expr",
        "googleapis/api/monitoredres",
        "protobuf/field_mask",
        "googleapis/api/annotations",
        "googleapis/api",
    ],
    deps = [
        ":google.golang.org.protobuf",
        ":github.com.golang.protobuf",
    ],
)

go_module(
    name = "github.com.hashicorp.raft-autopilot",
    module = "github.com/hashicorp/raft-autopilot",
    version = "v0.1.3",
    deps = [
        ":github.com.hashicorp.raft",
        ":golang.org.x.sync",
        ":github.com.hashicorp.go-hclog",
    ],
)

go_module(
    name = "github.com.joyent.triton-go",
    module = "github.com/joyent/triton-go",
    version = "v1.7.1-0.20200416154420-6801d15b779f",
    install = [
        ".",
        "authentication",
        "client",
        "compute",
        "storage",
        "errors",
    ],
    deps = [
        ":github.com.pkg.errors",
        ":golang.org.x.crypto",
    ],
)

go_module(
    name = "github.com.imdario.mergo",
    module = "github.com/imdario/mergo",
    version = "v0.3.11",
)

go_module(
    name = "github.com.go-asn1-ber.asn1-ber",
    module = "github.com/go-asn1-ber/asn1-ber",
    version = "v1.5.1",
)

go_module(
    name = "github.com.googleapis.gax-go.v2",
    module = "github.com/googleapis/gax-go/v2",
    version = "v2.0.5",
    deps = [
        ":google.golang.org.grpc",
    ],
)

go_module(
    name = "github.com.hashicorp.raft",
    module = "github.com/hashicorp/raft",
    version = "v1.3.0",
    deps = [
        ":github.com.hashicorp.go-hclog",
        ":github.com.hashicorp.go-msgpack",
        ":github.com.armon.go-metrics",
    ],
)

go_module(
    name = "github.com.hashicorp.vault-plugin-auth-oci",
    module = "github.com/hashicorp/vault-plugin-auth-oci",
    version = "v0.8.0",
    deps = [
        ":github.com.pkg.errors",
        ":github.com.hashicorp.vault.api",
        ":github.com.hashicorp.vault.sdk_1",
        ":github.com.oracle.oci-go-sdk",
        ":github.com.hashicorp.errwrap",
        ":github.com.hashicorp.go-hclog",
    ],
)

go_module(
    name = "github.com.hashicorp.go-uuid",
    module = "github.com/hashicorp/go-uuid",
    version = "v1.0.2",
)

go_module(
    name = "golang.org.x.sync",
    module = "golang.org/x/sync",
    version = "v0.0.0-20210220032951-036812b2e83c",
    install = [
        "errgroup",
        "semaphore",
    ],
)

go_module(
    name = "github.com.nwaples.rardecode",
    module = "github.com/nwaples/rardecode",
    version = "v1.1.0",
)

go_module(
    name = "github.com.elazarl.go-bindata-assetfs",
    module = "github.com/elazarl/go-bindata-assetfs",
    version = "v1.0.1-0.20200509193318-234c15e7648f",
)

go_module(
    name = "github.com.json-iterator.go",
    module = "github.com/json-iterator/go",
    version = "v1.1.11",
    deps = [
        ":github.com.modern-go.concurrent",
        ":github.com.modern-go.reflect2",
    ],
)

go_module(
    name = "github.com.Azure.go-autorest.autorest",
    module = "github.com/Azure/go-autorest/autorest",
    version = "v0.11.18",
    install = [
        ".",
        "azure",
    ],
    deps = [
        ":github.com.Azure.go-autorest.tracing",
        ":github.com.Azure.go-autorest.logger",
        ":github.com.Azure.go-autorest.autorest.adal",
    ],
)

go_module(
    name = "github.com.hashicorp.nomad.api",
    module = "github.com/hashicorp/nomad/api",
    version = "v0.0.0-20191220223628-edc62acd919d",
    install = [
        ".",
        "contexts",
    ],
    deps = [
        ":github.com.gorhill.cronexpr",
        ":github.com.hashicorp.go-rootcerts",
        ":github.com.hashicorp.go-cleanhttp",
        ":github.com.gorilla.websocket",
    ],
)

go_module(
    name = "github.com.gammazero.workerpool",
    module = "github.com/gammazero/workerpool",
    version = "v0.0.0-20190406235159-88d534f22b56",
    deps = [
        ":github.com.gammazero.deque",
    ],
)

go_module(
    name = "github.com.mitchellh.go-testing-interface",
    module = "github.com/mitchellh/go-testing-interface",
    version = "v1.14.0",
)

go_module(
    name = "github.com.chrismalek.oktasdk-go",
    module = "github.com/chrismalek/oktasdk-go",
    version = "v0.0.0-20181212195951-3430665dfaa0",
    install = [
        "okta",
    ],
    deps = [
        ":github.com.google.go-querystring",
    ],
)

go_module(
    name = "github.com.denverdino.aliyungo",
    module = "github.com/denverdino/aliyungo",
    version = "v0.0.0-20170926055100-d3308649c661",
    install = [
        "common",
        "util",
        "ecs",
    ],
)

go_module(
    name = "github.com.bgentry.speakeasy",
    module = "github.com/bgentry/speakeasy",
    version = "v0.1.0",
)

go_module(
    name = "github.com.Masterminds.sprig",
    module = "github.com/Masterminds/sprig",
    version = "v2.22.0+incompatible",
    deps = [
        ":github.com.google.uuid",
        ":golang.org.x.crypto",
        ":github.com.Masterminds.semver",
        ":github.com.huandu.xstrings",
        ":github.com.imdario.mergo",
        ":github.com.Masterminds.goutils",
        ":github.com.mitchellh.copystructure",
    ],
)

go_module(
    name = "github.com.sethvargo.go-limiter",
    module = "github.com/sethvargo/go-limiter",
    version = "v0.3.0",
    install = [
        ".",
        "memorystore",
        "internal/fasttime",
        "httplimit",
    ],
)

go_module(
    name = "github.com.hashicorp.vault-plugin-secrets-kv",
    module = "github.com/hashicorp/vault-plugin-secrets-kv",
    version = "v0.9.0",
    deps = [
        ":github.com.golang.protobuf",
        ":github.com.hashicorp.vault.sdk_1",
        ":github.com.mitchellh.mapstructure",
    ],
)

go_module(
    name = "gopkg.in.inf.v0",
    module = "gopkg.in/inf.v0",
    version = "v0.9.1",
)

go_module(
    name = "github.com.rboyer.safeio",
    module = "github.com/rboyer/safeio",
    version = "v0.2.1",
)

go_module(
    name = "github.com.okta.okta-sdk-golang.v2",
    module = "github.com/okta/okta-sdk-golang/v2",
    version = "v2.0.0",
    install = [
        "okta",
        "okta/query",
        "okta/cache",
    ],
    deps = [
        ":gopkg.in.square.go-jose.v2",
        ":github.com.go-yaml.yaml",
        ":github.com.kelseyhightower.envconfig",
        ":github.com.okta.okta-sdk-golang",
        ":github.com.patrickmn.go-cache",
    ],
)

go_module(
    name = "github.com.linode.linodego",
    module = "github.com/linode/linodego",
    version = "v0.7.1",
    deps = [
        ":gopkg.in.resty.v1",
    ],
)

go_module(
    name = "github.com.mitchellh.reflectwalk",
    module = "github.com/mitchellh/reflectwalk",
    version = "v1.0.1",
)

go_module(
    name = "github.com.Azure.go-autorest.autorest.azure.cli",
    module = "github.com/Azure/go-autorest/autorest/azure/cli",
    version = "v0.4.2",
    deps = [
        ":github.com.Azure.go-autorest.autorest.adal",
        ":github.com.dimchansky.utfbom",
        ":github.com.Azure.go-autorest.autorest.date",
        ":github.com.mitchellh.go-homedir",
    ],
)

go_module(
    name = "github.com.gregjones.httpcache",
    module = "github.com/gregjones/httpcache",
    version = "v0.0.0-20180305231024-9cad4c3443a7",
    install = [
        ".",
        "diskcache",
    ],
    deps = [
        ":github.com.peterbourgon.diskv",
    ],
)

go_module(
    name = "github.com.google.flatbuffers",
    module = "github.com/google/flatbuffers",
    version = "v1.11.0",
    install = [
        "go",
    ],
)

go_module(
    name = "github.com.golang.snappy",
    module = "github.com/golang/snappy",
    version = "v0.0.4",
)

go_module(
    name = "github.com.oklog.run",
    module = "github.com/oklog/run",
    version = "v1.0.0",
)

go_module(
    name = "github.com.go-stack.stack",
    module = "github.com/go-stack/stack",
    version = "v1.8.0",
)
