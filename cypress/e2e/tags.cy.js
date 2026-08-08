describe('Tags Page Tests', () => {
    context('Tag page - Luther Pass', () => {
        beforeEach(() => {
            cy.visit('/tags/luther-pass/')
        })

        it('should display tag title', () => {
            cy.get('h1').should('be.visible')
            cy.get('h1').should('contain', 'Luther Pass')
        })

        it('should have working breadcrumb navigation', () => {
            cy.get('nav[aria-label="Breadcrumb"]').should('be.visible')
            
            // Test home breadcrumb link
            cy.get('nav[aria-label="Breadcrumb"] a[href="/"]').click()
            cy.url().should('eq', Cypress.config().baseUrl)
        })

        it('should have tags breadcrumb link', () => {
            cy.visit('/tags/luther-pass/')
            
            // Test tags index breadcrumb link
            cy.contains('nav[aria-label="Breadcrumb"] a', 'Tags').should('be.visible')
        })

        it('should have working global navigation link', () => {
            // Test Recent link from global nav
            cy.contains('nav a', 'Recent').click()
            cy.url().should('include', '/recent')
        })

        it('should display sign content', () => {
            // Tags pages show signs with that tag
            cy.get('a[href*="/sign/"]').should('exist')
        })

        it('should navigate to sign page from tag content', () => {
            // Click on a sign link
            cy.get('a[href*="/sign/"]').first().click()
            cy.url().should('include', '/sign/')
        })
    })

    context('Tag page - New Jersey Turnpike', () => {
        beforeEach(() => {
            cy.visit('/tags/new-jersey-turnpike/')
        })

        it('should display tag title', () => {
            cy.get('h1').should('be.visible')
            cy.get('h1').should('contain', 'New Jersey Turnpike')
        })

        it('should have sign content', () => {
            cy.get('a[href*="/sign/"]').should('exist')
        })

        it('should have breadcrumb navigation', () => {
            cy.get('nav[aria-label="Breadcrumb"]').should('be.visible')
        })
    })

    context('Tag page with pagination', () => {
        it('should handle pagination if tag has many signs', () => {
            // Visit a tag that might have multiple pages
            cy.visit('/tags/the-dalles-dam/')
            
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

    context('Tag navigation', () => {
        it('should show related tags if available', () => {
            cy.visit('/tags/luther-pass/')
            
            // Some tag pages may show related tags
            cy.get('body').then($body => {
                if ($body.find('a[href*="/tags/"]').length > 1) {
                    // Find another tag link (not the current one)
                    cy.get('a[href*="/tags/"]').eq(1).then($link => {
                        const href = $link.attr('href')
                        if (href && href !== '/tags/luther-pass/') {
                            cy.wrap($link).click()
                            cy.url().should('include', '/tags/')
                            cy.get('h1').should('be.visible')
                        }
                    })
                }
            })
        })
    })
})
