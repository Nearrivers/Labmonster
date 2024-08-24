import { describe, expect, test } from 'vitest'
import { mount } from '@vue/test-utils'
import DeleteFileDialog from '../DeleteFileDialog.vue'

function setup() {
  const wrapper = mount(DeleteFileDialog, {
    data() {
      return {
        isDialogOpen: true
      }
    }
  })

  return wrapper
}

describe('DeleFileDialog', () => {
  test("open dialog", async () => {
    const wrapper = setup()
    wrapper.vm.openDialog("test")

    const dialog = wrapper.find('[data-test="dialog"]')
    expect(dialog.exists()).toBe(true)
  })

  // test("cancel dialog", async () => {
  //   const wrapper = setup()
  //   wrapper.vm.openDialog("test")

  //   const dialog = wrapper.get('[data-test="dialog"]')
  //   const cancelBtn = dialog.get('[data-test="cancel"]')
  // })
})