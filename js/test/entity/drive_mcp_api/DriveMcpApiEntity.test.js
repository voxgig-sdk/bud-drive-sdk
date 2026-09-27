
const envlocal = __dirname + '/../../../.env.local'
require('../../utility').loadEnvLocal(envlocal)

const Path = require('node:path')
const Fs = require('node:fs')

const { test, describe, afterEach } = require('node:test')
const assert = require('node:assert')
const { createLiveTransport } = require('../../live-runner')
const { runLiveEntity } = require('../../live-entity')


const { BudDriveSDK, BaseFeature, stdutil, config } = require('../../..')

const {
  envOverride,
  liveClientOptions,
  liveDelay,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
} = require('../../utility')


describe('DriveMcpApiEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when BUD_DRIVE_TEST_LIVE=TRUE.
  afterEach(liveDelay('BUD_DRIVE_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = BudDriveSDK.test()
    const ent = testsdk.DriveMcpApi()
    assert(null != ent)
  })


  test('basic', async (t) => {

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"error":{"a":true,"h":"Error","n":"error","r":false,"t":"`$OBJECT`","key$":"error","index$":0},"id":{"a":true,"h":"Id","n":"id","op":{"create":{"req":false,"type":"`$STRING`"}},"r":true,"t":"`$STRING`","union":{"branches":2,"count":1,"depth":0},"key$":"id","index$":1},"jsonrpc":{"a":true,"h":"Jsonrpc","n":"jsonrpc","r":true,"t":"`$STRING`","key$":"jsonrpc","index$":2},"method":{"a":true,"h":"Method","n":"method","r":true,"t":"`$STRING`","key$":"method","index$":3},"params":{"a":true,"h":"Params","n":"params","r":false,"t":"`$OBJECT`","key$":"params","index$":4},"result":{"a":true,"h":"Result","n":"result","r":false,"t":"`$OBJECT`","key$":"result","index$":5}},"id":{"field":"id","name":"id"},"name":"drive_mcp_api","op":{"create":{"input":"data","name":"create","points":[{"a":true,"co":{"id":"POST /drive-mcp/v1/campaigns","source":"openapi3","version":2},"g":{"header":[{"a":true,"ex":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","k":"header","n":"x_client_id","or":"x_client_id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"POST","o":"/drive-mcp/v1/campaigns","q":{"exist":["x_client_id"]},"r":{},"s":[{"lit":"drive-mcp"},{"lit":"v1"},{"lit":"campaigns"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0},{"a":true,"co":{"id":"POST /drive-mcp/v1/copilot","source":"openapi3","version":2},"g":{"header":[{"a":true,"ex":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","k":"header","n":"x_client_id","or":"x_client_id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"POST","o":"/drive-mcp/v1/copilot","q":{"exist":["x_client_id"]},"r":{},"s":[{"lit":"drive-mcp"},{"lit":"v1"},{"lit":"copilot"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":1},{"a":true,"co":{"id":"POST /drive-mcp/v1/segments","source":"openapi3","version":2},"g":{"header":[{"a":true,"ex":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","k":"header","n":"x_client_id","or":"x_client_id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"POST","o":"/drive-mcp/v1/segments","q":{"exist":["x_client_id"]},"r":{},"s":[{"lit":"drive-mcp"},{"lit":"v1"},{"lit":"segments"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":2}],"key$":"create"}},"relations":{"ancestors":[]},"key$":"drive_mcp_api","name__orig":"drive_mcp_api","Name":"DriveMcpApi","name_":"drive_mcp_api","name-":"drive-mcp-api","NAME":"DRIVE_MCP_API","index$":1}, {"active":true,"entity":"drive_mcp_api","key$":"BasicDriveMcpApiFlow","kind":"basic","name":"BasicDriveMcpApiFlow","param":{},"step":[{"a":true,"d":{},"i":{"ref":"drive_mcp_api_ref01"},"m":{},"o":"create","s":[],"v":[],"index$":0}]}, 'DriveMcpApi', {"POST /drive-mcp/v1/campaigns":{"protocol":"http","requestBody":{"required":true,"content":{"application/json":{"schema":{"title":"MCP JSON-RPC Request","description":"A [JSON-RPC 2.0](https://www.jsonrpc.org/specification) request, per the Model Context Protocol specification.","type":"object","required":["jsonrpc","method"],"properties":{"jsonrpc":{"type":"string","enum":["2.0"],"key$":"jsonrpc"},"id":{"oneOf":[{"type":"string"},{"type":"integer"}],"key$":"id"},"method":{"type":"string","key$":"method"},"params":{"type":"object","key$":"params"}},"example":{"jsonrpc":"2.0","id":1,"method":"tools/list"},"x-ref":"#/paths/~1drive-mcp~1v1~1copilot/post/requestBody/content/application~1json/schema","index$":1}}}},"parameters":[{"in":"header","name":"X-Client-Id","schema":{"type":"string"},"required":true,"description":"The API Client Identifier (Service Application Identifier).","example":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","x-ref":"#/paths/~1drive-api~1v1~1tags/get/parameters/0","index$":0}]},"POST /drive-mcp/v1/copilot":{"protocol":"http","requestBody":{"required":true,"content":{"application/json":{"schema":{"title":"MCP JSON-RPC Request","description":"A [JSON-RPC 2.0](https://www.jsonrpc.org/specification) request, per the Model Context Protocol specification.","type":"object","required":["jsonrpc","method"],"properties":{"jsonrpc":{"type":"string","enum":["2.0"],"key$":"jsonrpc"},"id":{"oneOf":[{"type":"string"},{"type":"integer"}],"key$":"id"},"method":{"type":"string","key$":"method"},"params":{"type":"object","key$":"params"}},"example":{"jsonrpc":"2.0","id":1,"method":"tools/list"},"index$":1}}}},"parameters":[{"in":"header","name":"X-Client-Id","schema":{"type":"string"},"required":true,"description":"The API Client Identifier (Service Application Identifier).","example":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","x-ref":"#/paths/~1drive-api~1v1~1tags/get/parameters/0","index$":0}]},"POST /drive-mcp/v1/segments":{"protocol":"http","requestBody":{"required":true,"content":{"application/json":{"schema":{"title":"MCP JSON-RPC Request","description":"A [JSON-RPC 2.0](https://www.jsonrpc.org/specification) request, per the Model Context Protocol specification.","type":"object","required":["jsonrpc","method"],"properties":{"jsonrpc":{"type":"string","enum":["2.0"],"key$":"jsonrpc"},"id":{"oneOf":[{"type":"string"},{"type":"integer"}],"key$":"id"},"method":{"type":"string","key$":"method"},"params":{"type":"object","key$":"params"}},"example":{"jsonrpc":"2.0","id":1,"method":"tools/list"},"x-ref":"#/paths/~1drive-mcp~1v1~1copilot/post/requestBody/content/application~1json/schema","index$":1}}}},"parameters":[{"in":"header","name":"X-Client-Id","schema":{"type":"string"},"required":true,"description":"The API Client Identifier (Service Application Identifier).","example":"8711bbef-b357-4c2d-97ca-0d9df4206e9a","x-ref":"#/paths/~1drive-api~1v1~1tags/get/parameters/0","index$":0}]}})
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const drive_mcp_api_ref01_ent = client.DriveMcpApi()
    let drive_mcp_api_ref01_data = setup.data.new.drive_mcp_api['drive_mcp_api_ref01']

    drive_mcp_api_ref01_data = (await drive_mcp_api_ref01_ent.create(drive_mcp_api_ref01_data)).data()
    assert(null != drive_mcp_api_ref01_data.id)


  })
})



function basicSetup(extra) {
  // TODO: fix test def options
  const options = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname,
      '../../../../.sdk/test/entity/drive_mcp_api/DriveMcpApiTestData.json')

  // TODO: file ready util needed?
  const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8')

  // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
  const entityData = JSON.parse(entityDataSource)

  options.entity = entityData.existing

  let client = BudDriveSDK.test(options, extra)
  const struct = client.utility().struct
  const merge = struct.merge
  const transform = struct.transform

  let idmap = transform(
    ['drive_mcp_api01','drive_mcp_api02','drive_mcp_api03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'BUD_DRIVE_TEST_DRIVE_MCP_API_ENTID': idmap,
    'BUD_DRIVE_TEST_LIVE': 'FALSE',
    'BUD_DRIVE_TEST_EXPLAIN': 'FALSE',
    'BUD_DRIVE_APIKEY': '',
  })

  idmap = env['BUD_DRIVE_TEST_DRIVE_MCP_API_ENTID']

  const live = 'TRUE' === env.BUD_DRIVE_TEST_LIVE
  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['BUD_DRIVE_TEST_DRIVE_MCP_API_ENTID']
    idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {}
    if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
      throw new Error('Live ENTID must be a JSON object')
    }
    client = new BudDriveSDK(merge([
      // FIRST, so the generated fields below win: sdk-test-control.json's
      // test.client.options adds to the live client, it does not redirect it.
      liveClientOptions(),
      {
        apikey: env.BUD_DRIVE_APIKEY,
      },
      // 'extra || {}', not a bare 'extra': struct.merge returns UNDEFINED when
      // the last entry is undefined, and basicSetup is normally called with no
      // argument at all - so a bare 'extra' silently discarded the apikey and
      // server values above and handed the SDK undefined.
      extra || {},
      { system: { fetch: transport.fetch } }
    ]))
  }

  const setup = {
    idmap,
    env,
    options,
    client,
    struct,
    data: entityData,
    explain: 'TRUE' === env.BUD_DRIVE_TEST_EXPLAIN,
    live,
    transport,
    now: Date.now(),
  }

  return setup
}
  
