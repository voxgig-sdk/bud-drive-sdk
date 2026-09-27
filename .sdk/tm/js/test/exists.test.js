
const { test, describe } = require('node:test')
const { equal } = require('node:assert')


const { BudDriveSDK } = require('..')


describe('exists', async () => {

  test('test-mode', async () => {
    const testsdk = await BudDriveSDK.test()
    equal(null !== testsdk, true)
  })

})
