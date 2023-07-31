import TodoItem from '../TodoItem.vue'
import type { Todo } from '@/App.vue'

describe('Todo Item', () => {
  const todo: Todo = {
    id: 1,
    name: 'Test',
    finished: false,
    editFlag: false
  }

  beforeEach(() => {
    cy.mount(TodoItem, {
      propsData: {
        todo
      }
    })
  })

  it('should check the checkbox correctly', () => {
    cy.get('[id=item-checkbox]').should('not.be.checked')
    cy.get('[id=item-checkbox').click()
    cy.get('[id=item-checkbox').should('be.checked')
  })

  it('should have delete button', () => {
    cy.get('.delete').should('exist')
    cy.get('.delete').click()
  })

  it('should convert to save button after pressing edit', () => {
    cy.get('.edit').should('exist')
    cy.get('.edit').click()

    cy.get('#input').should('exist')
    cy.get('.save').should('exist')
  })
})
