describe('Category Page Tests', () => {
    context('Category page - Alabama Interstate Shield with State Name', () => {
        beforeEach(() => {
            cy.visit('/categories/alabama_shield-interstate-state-name/')
        })

        it('should display category title with state and category name', () => {
            cy.get('h1').should('be.visible')
            cy.get('h1').should('contain', 'Alabama')
            cy.get('h1').should('contain', 'Interstate Shield')
        })

        it('should have working breadcrumb navigation', () => {
            cy.get('nav[aria-label="Breadcrumb"]').should('be.visible')
            
            // Test state breadcrumb link
            cy.contains('nav[aria-label="Breadcrumb"] a', 'Alabama').click()
            cy.url().should('include', '/state/alabama')
        })

        it('should navigate to country from breadcrumb', () => {
            cy.visit('/categories/alabama_shield-interstate-state-name/')
            
            // Test country breadcrumb link
            cy.contains('nav[aria-label="Breadcrumb"] a', 'United States').click()
            cy.url().should('include', '/country/united-states')
        })

        it('should have working global navigation link', () => {
            cy.visit('/categories/alabama_shield-interstate-state-name/')
            
            // Test Home link from global nav
            cy.contains('nav a', 'Home').click()
            cy.url().should('eq', Cypress.config().baseUrl)
        })

        it('should display sign content', () => {
            // Category pages show filtered signs
            cy.get('a[href*="/sign/"]').should('exist')
        })

        it('should navigate to sign page', () => {
            // Click on a sign link
            cy.get('a[href*="/sign/"]').first().click()
            cy.url().should('include', '/sign/')
        })
    })

    context('Category page - Another category', () => {
        beforeEach(() => {
            cy.visit('/categories/ohio_bridge-named/')
        })

        it('should display category title', () => {
            cy.get('h1').should('be.visible')
            cy.get('h1').should('contain', 'Ohio')
        })

        it('should have breadcrumb with state link', () => {
            cy.contains('nav[aria-label="Breadcrumb"] a', 'Ohio').should('be.visible')
        })

        it('should have sign content', () => {
            cy.get('a[href*="/sign/"]').should('exist')
        })
    })

    context('Category page with pagination', () => {
        it('should handle pagination if category has many signs', () => {
            cy.visit('/categories/alabama_shield-interstate-state-name/')
            
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

    context('Category navigation hierarchy', () => {
        it('should maintain proper geographic context in breadcrumb', () => {
            cy.visit('/categories/alabama_shield-interstate-state-name/')
            
            // Verify full breadcrumb hierarchy
            cy.get('nav[aria-label="Breadcrumb"]').should('be.visible')
            cy.contains('nav[aria-label="Breadcrumb"] a', 'Home').should('exist')
            cy.contains('nav[aria-label="Breadcrumb"] a', 'United States').should('exist')
            cy.contains('nav[aria-label="Breadcrumb"] a', 'Alabama').should('exist')
        })
    })
})
