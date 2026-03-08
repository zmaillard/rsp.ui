describe('County Page Tests', () => {
    
    context('Single Locality County - Wahkiakum', () => {
        beforeEach(() => {
            cy.visit('/county/washington_wahkiakum-county/')
        })

        it('should display the page title with county and state name', () => {
            cy.get('h1').should('contain', 'Wahkiakum County')
            cy.get('h1').should('contain', 'Washington')
        })

        it('should display correct breadcrumb navigation', () => {
            cy.get('[data-cy="breadcrumb"]').should('be.visible')
            cy.get('[data-cy="breadcrumb-home"]').should('be.visible')
            cy.get('[data-cy="breadcrumb-country"]').should('be.visible').and('contain', 'United States')
            cy.get('[data-cy="breadcrumb-state"]').should('be.visible').and('contain', 'Washington')
            cy.get('[data-cy="breadcrumb-current"]').should('be.visible').and('contain', 'Wahkiakum County')
        })

        it('should have working breadcrumb links', () => {
            // Test home link
            cy.get('[data-cy="breadcrumb-home"] a').click()
            cy.url().should('eq', Cypress.config().baseUrl)
            cy.go('back')

            // Test country link
            cy.get('[data-cy="breadcrumb-country"] a').click()
            cy.url().should('include', '/country/')
            cy.go('back')

            // Test state link
            cy.get('[data-cy="breadcrumb-state"] a').click()
            cy.url().should('include', '/state/washington')
        })

        it('should display the "Localities:" header', () => {
            cy.get('[data-cy="localities-header"]').should('be.visible')
            cy.get('[data-cy="localities-header"]').should('contain', 'Localities:')
        })

        it('should display exactly one locality link with no pipe delimiters', () => {
            cy.get('[data-cy="locality-link"]').should('have.length', 1)
            cy.get('[data-cy="locality-delimiter"]').should('not.exist')
        })

        it('should NOT contain any pipe characters in localities section', () => {
            cy.get('[data-cy="localities-header"]').invoke('text').should('not.contain', '|')
        })

        it('should display place section with correct heading', () => {
            cy.get('[data-cy="place-heading"]').should('be.visible')
            cy.get('[data-cy="place-heading"]').should('have.length.at.least', 1)
        })

        it('should display sign tiles', () => {
            cy.get('[data-cy="sign-tile"]').should('be.visible')
            cy.get('[data-cy="sign-tile"]').should('have.length.at.least', 1)
        })

        it('should have clickable sign tiles that navigate to sign pages', () => {
            cy.get('[data-cy="sign-tile"]').first().find('a').first().click()
            cy.url().should('include', '/sign/')
        })

        it('should have working random sign link', () => {
            cy.get('[data-cy="random-sign-link"]').should('be.visible')
            cy.get('[data-cy="random-sign-link"]').should('have.attr', 'href')
                .and('include', 'washington_wahkiakum-county')
        })

        it('should navigate locality anchor link to correct place section', () => {
            cy.get('[data-cy="locality-link"]').first().then($link => {
                const href = $link.attr('href')
                const targetId = href.replace('#', '')
                
                cy.get('[data-cy="locality-link"]').first().click()
                
                // Verify the hash is in the URL
                cy.hash().should('eq', `#${targetId}`)
                
                // Verify the target section exists
                cy.get(`#${targetId}`).should('exist')
            })
        })
    })

    context('Two Localities County - Clallam', () => {
        beforeEach(() => {
            cy.visit('/county/washington_clallam-county/')
        })

        it('should display the page title correctly', () => {
            cy.get('h1').should('contain', 'Clallam County')
            cy.get('h1').should('contain', 'Washington')
        })

        it('should display multiple locality links', () => {
            cy.get('[data-cy="locality-link"]').should('have.length.at.least', 2)
        })

        it('should have exactly one pipe delimiter between two localities', () => {
            cy.get('[data-cy="locality-link"]').its('length').then((count) => {
                // For n localities, there should be n-1 pipes
                cy.get('[data-cy="locality-delimiter"]').should('have.length', count - 1)
            })
        })

        it('should format pipes without spaces', () => {
            // Get the text content of the localities section
            cy.get('[data-cy="localities-header"]').invoke('text').then((text) => {
                // Should not have spaces around pipes
                expect(text).to.not.match(/ \|/)
                expect(text).to.not.match(/\| /)
            })
        })

        it('should not have leading or trailing pipe delimiters', () => {
            cy.get('[data-cy="localities-header"]').invoke('text').then((text) => {
                // Extract just the locality links portion (after "Localities:")
                const localitiesText = text.replace('Localities:', '').trim()
                
                // Should not start or end with pipe
                expect(localitiesText).to.not.match(/^\|/)
                expect(localitiesText).to.not.match(/\|$/)
            })
        })

        it('should navigate to correct place sections when locality links clicked', () => {
            cy.get('[data-cy="locality-link"]').first().then($link => {
                const href = $link.attr('href')
                const targetId = href.replace('#', '')
                
                cy.get('[data-cy="locality-link"]').first().click()
                cy.hash().should('eq', `#${targetId}`)
                cy.get(`#${targetId}`).should('exist')
            })
        })

        it('should display multiple place sections', () => {
            cy.get('[data-cy="place-heading"]').should('have.length.at.least', 2)
        })

        it('should display sign tiles in each place section', () => {
            cy.get('[data-cy="sign-tile"]').should('be.visible')
            cy.get('[data-cy="sign-tile"]').should('have.length.at.least', 2)
        })
    })

    context('Three Localities County - Fairfax', () => {
        beforeEach(() => {
            cy.visit('/county/virginia_fairfax-county/')
        })

        it('should display the page title correctly', () => {
            cy.get('h1').should('contain', 'Fairfax County')
            cy.get('h1').should('contain', 'Virginia')
        })

        it('should display exactly three locality links', () => {
            cy.get('[data-cy="locality-link"]').should('have.length', 3)
        })

        it('should have exactly two pipe delimiters for three localities', () => {
            cy.get('[data-cy="locality-delimiter"]').should('have.length', 2)
        })

        it('should format as "LocalityA|LocalityB|LocalityC" with no spaces', () => {
            cy.get('[data-cy="localities-header"]').invoke('text').then((text) => {
                // Should have exactly 2 pipes
                const pipeCount = (text.match(/\|/g) || []).length
                expect(pipeCount).to.equal(2)
                
                // Should not have spaces around pipes
                expect(text).to.not.match(/ \|/)
                expect(text).to.not.match(/\| /)
            })
        })

        it('should not have leading or trailing pipes', () => {
            cy.get('[data-cy="localities-header"]').invoke('text').then((text) => {
                const localitiesText = text.replace('Localities:', '').trim()
                
                // Remove all whitespace to check pipe positions
                const normalized = localitiesText.replace(/\s+/g, '')
                
                // Should not start or end with pipe
                expect(normalized).to.not.match(/^\|/)
                expect(normalized).to.not.match(/\|$/)
            })
        })

        it('should display three place sections with headings', () => {
            cy.get('[data-cy="place-heading"]').should('have.length', 3)
        })

        it('should have all locality anchor links working', () => {
            cy.get('[data-cy="locality-link"]').each(($link) => {
                const href = $link.attr('href')
                const targetId = href.replace('#', '')
                
                // Verify the target section exists
                cy.get(`#${targetId}`).should('exist')
            })
        })
    })

    context('County with "No Place Associated" signs', () => {
        // Using Clallam County as it's large enough to potentially have unassociated signs
        beforeEach(() => {
            cy.visit('/county/washington_clallam-county/')
        })

        it('should display "No Place Associated" section if signs without place exist', () => {
            cy.get('body').then($body => {
                if ($body.find('[data-cy="no-place-heading"]').length > 0) {
                    cy.get('[data-cy="no-place-heading"]').should('be.visible')
                    cy.get('[data-cy="no-place-heading"]').should('contain', 'No Place Associated')
                    
                    // Should have sign tiles in the no-place section
                    cy.get('[data-cy="no-place-heading"]').parent().find('[data-cy="sign-tile"]')
                        .should('have.length.at.least', 1)
                }
            })
        })
    })

    context('Common County Page Functionality', () => {
        beforeEach(() => {
            cy.visit('/county/washington_clallam-county/')
        })

        it('should have complete breadcrumb navigation hierarchy', () => {
            cy.get('[data-cy="breadcrumb"]').should('be.visible')
            
            // Verify all breadcrumb levels exist
            cy.get('[data-cy="breadcrumb-home"]').should('exist')
            cy.get('[data-cy="breadcrumb-country"]').should('exist')
            cy.get('[data-cy="breadcrumb-state"]').should('exist')
            cy.get('[data-cy="breadcrumb-current"]').should('exist')
        })

        it('should display "Localities:" header even if no localities', () => {
            cy.get('[data-cy="localities-header"]').should('be.visible')
            cy.get('[data-cy="localities-header"]').should('contain', 'Localities:')
        })

        it('should have valid random sign link URL structure', () => {
            cy.get('[data-cy="random-sign-link"]').should('have.attr', 'href')
                .and('match', /\/statesubdivision\/\w+_[\w-]+/)
        })

        it('should display sign tiles with proper structure', () => {
            cy.get('[data-cy="sign-tile"]').first().within(() => {
                cy.get('img').should('be.visible')
                cy.get('a').should('have.attr', 'href').and('include', '/sign/')
            })
        })

        it('should have place headings that are links to place pages', () => {
            cy.get('[data-cy="place-heading"] a').first().should('exist')
            cy.get('[data-cy="place-heading"] a').first().should('have.attr', 'href')
                .and('include', '/place/')
        })
    })

    context('Pipe Delimiter Edge Cases', () => {
        it('should correctly handle single locality (no pipes) - Wahkiakum', () => {
            cy.visit('/county/washington_wahkiakum-county/')
            
            cy.get('[data-cy="locality-link"]').should('have.length', 1)
            cy.get('[data-cy="locality-delimiter"]').should('not.exist')
            
            // Verify no pipe character anywhere in localities section
            cy.get('[data-cy="localities-header"]').invoke('text')
                .should('not.contain', '|')
        })

        it('should correctly handle two localities (one pipe) - Clallam', () => {
            cy.visit('/county/washington_clallam-county/')
            
            cy.get('[data-cy="locality-link"]').its('length').then((count) => {
                if (count === 2) {
                    cy.get('[data-cy="locality-delimiter"]').should('have.length', 1)
                    
                    // Verify format: "A|B"
                    cy.get('[data-cy="localities-header"]').invoke('text').then((text) => {
                        const pipeCount = (text.match(/\|/g) || []).length
                        expect(pipeCount).to.equal(1)
                    })
                }
            })
        })

        it('should correctly handle three localities (two pipes) - Fairfax', () => {
            cy.visit('/county/virginia_fairfax-county/')
            
            cy.get('[data-cy="locality-link"]').should('have.length', 3)
            cy.get('[data-cy="locality-delimiter"]').should('have.length', 2)
            
            // Verify format: "A|B|C"
            cy.get('[data-cy="localities-header"]').invoke('text').then((text) => {
                const pipeCount = (text.match(/\|/g) || []).length
                expect(pipeCount).to.equal(2)
            })
        })

        it('should verify mathematical relationship: pipes = localities - 1', () => {
            cy.visit('/county/virginia_fairfax-county/')
            
            cy.get('[data-cy="locality-link"]').its('length').then((localityCount) => {
                cy.get('[data-cy="locality-delimiter"]').its('length').then((pipeCount) => {
                    expect(pipeCount).to.equal(localityCount - 1)
                })
            })
        })

        it('should ensure pipes are direct siblings between links', () => {
            cy.visit('/county/virginia_fairfax-county/')
            
            // Verify the structure: link, pipe, link, pipe, link
            cy.get('[data-cy="localities-header"]').within(() => {
                cy.get('[data-cy="locality-link"]').first().next('[data-cy="locality-delimiter"]')
                    .should('exist')
            })
        })
    })

    context('Accessibility and Content Validation', () => {
        beforeEach(() => {
            cy.visit('/county/washington_clallam-county/')
        })

        it('should have proper heading hierarchy', () => {
            cy.get('h1').should('have.length', 1)
            cy.get('[data-cy="localities-header"]').should('match', 'h5')
            cy.get('[data-cy="place-heading"]').should('match', 'h4')
        })

        it('should have aria-label on breadcrumb navigation', () => {
            cy.get('[data-cy="breadcrumb"]').should('have.attr', 'aria-label', 'Breadcrumb')
        })

        it('should have aria-current on current breadcrumb item', () => {
            cy.get('[data-cy="breadcrumb-current"]').should('have.attr', 'aria-current', 'page')
        })

        it('should have images with alt text', () => {
            cy.get('[data-cy="sign-tile"] img').first().should('have.attr', 'alt')
        })

        it('should have all links with valid href attributes', () => {
            cy.get('[data-cy="locality-link"]').each(($link) => {
                cy.wrap($link).should('have.attr', 'href').and('not.be.empty')
            })
        })
    })
})
