import { describe, expect, test } from 'vitest'
import { mount } from '@vue/test-utils'
import DeleteFileDialog from '../DeleteFileDialog.vue'
import { nextTick } from 'vue'

describe('DeleFileDialog', () => {
  test("open dialog", async () => {
    const wrapper = mount(DeleteFileDialog, {
      data() {
        return {
          isDialogOpen: true
        }
      }
    })

    wrapper.vm.openDialog("test")
    await nextTick()

    // Ne fonctionnent pas à cause du warning "Extraneous props...."
    const dialog = wrapper.find('[data-test="dialog"]')
    expect(dialog.exists()).toBe(true)
    const description = wrapper.find('[data-test="description"]')
    expect(description.exists()).toBe(true)
  })
})