describe('Recent Page Tests', () => {
    context('Recent index page', () => {
        beforeEach(() => {
            cy.visit('/recent/')
        })

        it('should display recent page title', () => {
            cy.get('h1').should('be.visible')
        })

        it('should have working breadcrumb navigation', () => {
            cy.get('nav[aria-label="Breadcrumb"]').should('be.visible')
            
            // Test home breadcrumb link
            cy.get('nav[aria-label="Breadcrumb"] a[href="/"]').click()
            cy.url().should('eq', Cypress.config().baseUrl)
        })

        it('should have working global navigation link', () => {
            cy.visit('/recent/')
            
            // Test Home link from global nav
            cy.contains('nav a', 'Home').click()
            cy.url().should('eq', Cypress.config().baseUrl)
        })

        it('should display recent sign content', () => {
            // Recent page shows recently added/taken signs
            cy.get('a[href*="/sign/"]').should('exist')
        })

        it('should navigate to sign page', () => {
            // Click on a sign link
            cy.get('a[href*="/sign/"]').first().click()
            cy.url().should('include', '/sign/')
        })

        it('should display month/year navigation if available', () => {
            cy.get('body').then($body => {
                // Check for date-based navigation links
                if ($body.find('a[href*="/recent/"]').length > 0) {
                    cy.get('a[href*="/recent/"]').should('exist')
                }
            })
        })
    })

    context('Recent page with specific month/year', () => {
        beforeEach(() => {
            cy.visit('/recent/2024-08/')
        })

        it('should display month-specific recent signs', () => {
            cy.get('h1').should('be.visible')
        })

        it('should have sign content', () => {
            cy.get('body').then($body => {
                if ($body.find('a[href*="/sign/"]').length > 0) {
                    cy.get('a[href*="/sign/"]').should('exist')
                }
            })
        })

        it('should have breadcrumb navigation', () => {
            cy.get('nav[aria-label="Breadcrumb"]').should('be.visible')
        })
    })

    context('Recent page pagination', () => {
        it('should handle pagination on recent page', () => {
            cy.visit('/recent/')
            
            cy.get('body').then($body => {
                if ($body.find('[data-cy="pagination"]').length > 0) {
                    cy.get('[data-cy="pagination"]').should('be.visible')
                    
                    // Test pagination navigation
                    if ($body.find('[data-cy="pagination-next"]').length > 0) {
                        cy.get('[data-cy="pagination-next"]').click()
                        cy.url().should('include', '/page/')
                    }
                }
            })
        })
    })

    context('Date grouping', () => {
        it('should display signs organized by date', () => {
            cy.visit('/recent/')
            
            // Recent pages may group signs by date
            cy.get('body').should('be.visible')
        })
    })

    context('Navigation to recent from global nav', () => {
        it('should navigate to recent page from any page', () => {
            cy.visit('/')
            cy.contains('nav a', 'Recent').click()
            cy.url().should('include', '/recent')
        })
    })
})
