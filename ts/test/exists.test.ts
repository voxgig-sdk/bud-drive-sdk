
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { BudDriveSDK } from '..'


describe('exists', async () => {

  test('test-mode', () => {
    const testsdk = BudDriveSDK.test()
    equal(testsdk instanceof BudDriveSDK, true,
      'BudDriveSDK.test() must return a client synchronously')
  })

})
