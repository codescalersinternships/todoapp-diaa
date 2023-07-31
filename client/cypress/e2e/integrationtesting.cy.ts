describe('Todo Item', () => {
  let addedItems = 0
  beforeEach(() => {
    cy.request({
      method: 'POST',
      url: 'localhost:5000/todo',
      body: {
        name: 'Test todo'
      }
    })

    cy.visit('localhost:5173')
  })

  it('should set the task as finished and check the checkbox', () => {
    cy.get('[id=item-checkbox').first().click()
    cy.get('[id=item-checkbox]').first().should('be.checked')
    addedItems++
  })

  it('should be able to delete todo', () => {
    cy.get('.delete')
      .its('length')
      .then((len) => {
        cy.get('.delete').first().should('exist').click()
        cy.get('.delete')
          .its('length')
          .should('eq', len - 1)
      })
  })

  it('should edit name successfully', () => {
    cy.get('.edit').first().should('exist').click()
    cy.get('#input').clear()
    cy.get('#input').first().type('new value')

    cy.get('.save').first().should('exist').click()

    cy.get('.item-title').first().invoke('text').should('eq', 'new value')
    addedItems++
  })

  it('should add a new task sucessfully ', () => {
    cy.get('#title').should('exist').type('new task')
    cy.get('.add-btn').should('exist').click()
    cy.get('.item-title').first().invoke('text').should('eq', 'new task')
    cy.get('.delete').first().should('exist').click()
    addedItems++
  })

  after(() => {
    for (let i = 0; i < addedItems; i++) cy.get('.delete').first().click()
  })
})
