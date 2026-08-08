describe('About Page Tests', () => {
    context('About page', () => {
        beforeEach(() => {
            cy.visit('/about/')
        })

        it('should display about page title', () => {
            cy.get('h1').should('be.visible')
            cy.get('h1').should('contain', 'About')
        })

        it('should have working breadcrumb navigation', () => {
            cy.get('nav[aria-label="Breadcrumb"]').should('be.visible')
            
            // Test home breadcrumb link
            cy.get('nav[aria-label="Breadcrumb"] a[href="/"]').click()
            cy.url().should('eq', Cypress.config().baseUrl)
        })

        it('should have working global navigation link', () => {
            cy.visit('/about/')
            
            // Test Home link from global nav
            cy.contains('nav a', 'Home').click()
            cy.url().should('eq', Cypress.config().baseUrl)
        })

        it('should display content sections', () => {
            // About page should have informational content
            cy.get('body').should('be.visible')
            cy.get('section, article, div').should('exist')
        })

        it('should have proper external link attributes if external links present', () => {
            cy.get('body').then($body => {
                // Check if there are any external links
                const externalLinks = $body.find('a[href^="http"]:not([href*="roadsign.pictures"])')
                
                if (externalLinks.length > 0) {
                    // External links should have target="_blank" and rel="noopener"
                    cy.get('a[href^="http"]:not([href*="roadsign.pictures"])').each($link => {
                        cy.wrap($link).should('have.attr', 'target', '_blank')
                        cy.wrap($link).should('have.attr', 'rel').and('match', /noopener/)
                    })
                }
            })
        })
    })

    context('About page accessibility', () => {
        beforeEach(() => {
            cy.visit('/about/')
        })

        it('should have proper heading hierarchy', () => {
            cy.get('h1').should('have.length', 1)
        })

        it('should have readable content', () => {
            // About page should have text content
            cy.get('body').invoke('text').should('not.be.empty')
        })
    })

    context('Navigation to about page', () => {
        it('should be accessible from homepage if linked', () => {
            cy.visit('/')
            
            cy.get('body').then($body => {
                if ($body.find('a[href*="/about"]').length > 0) {
                    cy.get('a[href*="/about"]').first().click()
                    cy.url().should('include', '/about')
                }
            })
        })
    })
})
