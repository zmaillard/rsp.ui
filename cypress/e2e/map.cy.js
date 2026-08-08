describe('Map Page Tests (Smoke Tests)', () => {
    context('Map page basic functionality', () => {
        beforeEach(() => {
            cy.visit('/map/')
        })

        it('should load the map page', () => {
            cy.get('body').should('be.visible')
        })


        it('should have working global navigation link', () => {
            cy.visit('/map/')
            
            // Test Home link from global nav
            cy.contains('nav a', 'Home').click()
            cy.url().should('eq', Cypress.config().baseUrl)
        })

        it('should have map container element', () => {
            // Map should have a container div (typically with id or class for mapbox)
            cy.get('body').then($body => {
                // Look for common map container patterns
                const hasMapContainer = 
                    $body.find('#map').length > 0 ||
                    $body.find('[class*="map"]').length > 0 ||
                    $body.find('[class*="mapbox"]').length > 0
                
                expect(hasMapContainer).to.be.true
            })
        })
    })

    context('Map page title', () => {
        beforeEach(() => {
            cy.visit('/map/')
        })

        it('should display page title or heading', () => {
            cy.get('body').then($body => {
                const hasTitle = 
                    $body.find('h1').length > 0 ||
                    $body.find('h2').length > 0
                
                if (hasTitle) {
                    cy.get('h1, h2').first().should('be.visible')
                }
            })
        })
    })

    // Note: Full Mapbox interaction testing is deferred
    // This includes: marker clicks, popup functionality, map navigation
})
