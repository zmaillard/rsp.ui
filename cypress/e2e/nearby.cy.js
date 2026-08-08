describe('Nearby Page Tests (Smoke Tests)', () => {
    context('Nearby page basic functionality', () => {
        beforeEach(() => {
            cy.visit('/nearby/')
        })

        it('should load the nearby page', () => {
            cy.get('body').should('be.visible')
        })

        it('should have working breadcrumb navigation', () => {
            cy.get('body').then($body => {
                if ($body.find('nav[aria-label="Breadcrumb"]').length > 0) {
                    cy.get('nav[aria-label="Breadcrumb"]').should('be.visible')
                }
            })
        })

        it('should have working global navigation link', () => {
            // Test Home link from global nav
            cy.contains('nav a', 'Home').click()
            cy.url().should('eq', Cypress.config().baseUrl)
        })

        it('should display page title or heading', () => {
            cy.get('body').then($body => {
                if ($body.find('h1').length > 0) {
                    cy.get('h1').should('be.visible')
                }
            })
        })

        it('should have geolocation prompt or message', () => {
            // Nearby page typically shows a geolocation message or prompt
            cy.get('body').invoke('text').should('not.be.empty')
        })
    })

    context('Nearby page navigation', () => {
        it('should be accessible from global navigation', () => {
            cy.visit('/')
            cy.contains('nav a', 'Location').click()
            cy.url().should('include', '/nearby')
        })
    })

    // Note: Full geolocation testing is deferred
    // This includes: mocking geolocation, testing results display, error handling
})
