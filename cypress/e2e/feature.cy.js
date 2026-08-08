describe('Feature Page Tests', () => {
    context('Feature page - Exit 131', () => {
        beforeEach(() => {
            cy.visit('/feature/1000/')
        })

        it('should display feature title', () => {
            cy.get('h1').should('be.visible')
            cy.get('h1').should('contain', 'Exit')
        })

        it('should have working breadcrumb navigation', () => {
            cy.get('nav[aria-label="Breadcrumb"]').should('be.visible')
            
            // Test state breadcrumb link
            cy.contains('nav[aria-label="Breadcrumb"] a', 'Washington').click()
            cy.url().should('include', '/state/washington')
        })

        it('should display highway shield images in breadcrumb', () => {
            // Highway shields should be visible in breadcrumb
            cy.get('nav[aria-label="Breadcrumb"] img[alt*="-"]').should('exist')
        })

        it('should navigate to highway from breadcrumb shield', () => {
            // Click on a highway shield in breadcrumb
            cy.get('nav[aria-label="Breadcrumb"] a[href*="/highway/"]').first().click()
            cy.url().should('include', '/highway/')
        })

        it('should have working global navigation link', () => {
            cy.visit('/feature/1000/')
            
            // Test Map link from global nav
            cy.contains('nav a', 'Map').click()
            cy.url().should('include', '/map')
        })

        it('should display sign tiles and navigate to sign page', () => {
            cy.get('[data-cy="sign-tile"]').should('be.visible')
            cy.get('[data-cy="sign-tile"]').should('have.length.at.least', 1)
            
            // Test one sign tile link
            cy.get('[data-cy="sign-tile"]').first().find('a').first().click()
            cy.url().should('include', '/sign/')
        })

        it('should have previous or next feature navigation', () => {
            // Check if navigation arrows exist (may have one or both depending on position)
            cy.get('body').then($body => {
                const hasPrev = $body.find('a[href*="/feature/"]').length > 0
                const hasNav = $body.find('svg').length > 0
                
                // Should have some form of navigation present
                expect(hasPrev || hasNav).to.be.true
            })
        })
    })

    context('Feature page - Another exit', () => {
        beforeEach(() => {
            cy.visit('/feature/2500/')
        })

        it('should display feature title', () => {
            cy.get('h1').should('be.visible')
        })

        it('should have sign content', () => {
            cy.get('[data-cy="sign-tile"]').should('have.length.at.least', 1)
        })

        it('should have breadcrumb navigation', () => {
            cy.get('nav[aria-label="Breadcrumb"]').should('be.visible')
        })
    })

    context('Feature navigation', () => {
        it('should navigate between features using prev/next links', () => {
            cy.visit('/feature/1000/')
            
            // Try to find and click a feature navigation link
            cy.get('a[href*="/feature/"]').first().then($link => {
                const href = $link.attr('href')
                if (href && href !== '/feature/1000/') {
                    cy.wrap($link).click()
                    cy.url().should('include', '/feature/')
                    cy.get('h1').should('be.visible')
                }
            })
        })
    })
})
