describe('Highway Page Wikipedia Integration', () => {
    context('Highway with Wikipedia external link', () => {
        beforeEach(() => {
            // Intercept Wikipedia API calls
            cy.intercept('GET', '**/en.wikipedia.org/w/api.php*', (req) => {
                // Mock Wikipedia API response
                req.reply({
                    statusCode: 200,
                    body: {
                        query: {
                            pages: {
                                '12345': {
                                    pageid: 12345,
                                    title: 'State Route 195',
                                    extract: 'State Route 195 (SR 195) is a state highway in the U.S. state of Arizona. It connects Interstate 10 with the city of Phoenix.'
                                }
                            }
                        }
                    }
                })
            }).as('wikiApi')

            cy.visit('/highway/az195')
        })

        it('should display wiki-search element when external link exists without hash', () => {
            cy.get('#wiki-search').should('exist')
            cy.get('#wiki-search').should('have.attr', 'data-external-link')
        })

        it('should fetch Wikipedia content and display it', () => {
            // Wait for API call
            cy.wait('@wikiApi')

            // Verify content is displayed
            cy.get('#wiki-search')
                .should('contain', 'State Route 195')
                .should('contain', 'Source:')
                .should('contain', 'Wikipedia')
        })

        it('should include a link to Wikipedia source', () => {
            cy.wait('@wikiApi')

            cy.get('#wiki-search a')
                .should('have.attr', 'href')
                .and('include', 'wikipedia.org')

            cy.get('#wiki-search a')
                .should('have.attr', 'target', '_blank')
                .should('have.attr', 'rel', 'noopener noreferrer')
        })

        it('should display the full Wikipedia extract text', () => {
            cy.wait('@wikiApi')

            cy.get('#wiki-search')
                .should('contain', 'Interstate 10')
                .should('contain', 'Phoenix')
        })
    })

    context('Highway with external link containing hash', () => {
        beforeEach(() => {
            // Create a test page with a hash in the external link
            cy.intercept('GET', '**/en.wikipedia.org/w/api.php*').as('wikiApi')
        })

        it('should not fetch Wikipedia content when external link has hash', () => {
            // Visit a highway page where external link has a hash fragment
            cy.visit('/highway/tx37spur', {
                onBeforeLoad(win) {
                    // Mock the page to have an external link with hash
                    win.document.getElementById = () => ({
                        dataset: {
                            externalLink: 'https://en.wikipedia.org/wiki/Test#Section'
                        }
                    })
                }
            })

            // Should not make API call
            cy.get('@wikiApi.all').should('have.length', 0)
        })
    })

    context('Highway without external link', () => {
        beforeEach(() => {
            cy.visit('/highway/futurei26')
        })

        it('should not display wiki-search element when no external link exists', () => {
            cy.get('#wiki-search').should('not.exist')
        })
    })

    context('Wikipedia API error handling', () => {
        beforeEach(() => {
            cy.visit('/highway/az195')
        })

        it('should handle API failure gracefully', () => {
            cy.intercept('GET', '**/en.wikipedia.org/w/api.php*', {
                statusCode: 500,
                body: 'Internal Server Error'
            }).as('wikiApiFail')

            cy.wait('@wikiApiFail')

            // Element should still exist but remain empty
            cy.get('#wiki-search').should('exist')
            cy.get('#wiki-search').should('be.empty')
        })

        it('should handle missing page data', () => {
            cy.intercept('GET', '**/en.wikipedia.org/w/api.php*', {
                statusCode: 200,
                body: {
                    query: {
                        pages: {}
                    }
                }
            }).as('wikiApiEmpty')

            cy.wait('@wikiApiEmpty')

            cy.get('#wiki-search').should('be.empty')
        })

        it('should handle missing extract data', () => {
            cy.intercept('GET', '**/en.wikipedia.org/w/api.php*', {
                statusCode: 200,
                body: {
                    query: {
                        pages: {
                            '12345': {
                                pageid: 12345,
                                title: 'Test Highway'
                                // No extract field
                            }
                        }
                    }
                }
            }).as('wikiApiNoExtract')

            cy.wait('@wikiApiNoExtract')

            cy.get('#wiki-search').should('be.empty')
        })
    })

    context('Wikipedia API request format', () => {
        beforeEach(() => {
            cy.intercept('GET', '**/en.wikipedia.org/w/api.php*').as('wikiApi')
            cy.visit('/highway/az195')
        })

        it('should make API request with correct parameters', () => {
            cy.wait('@wikiApi').then((interception) => {
                const url = new URL(interception.request.url)
                const params = url.searchParams

                expect(params.get('action')).to.equal('query')
                expect(params.get('redirects')).to.equal('1')
                expect(params.get('explaintext')).to.equal('true')
                expect(params.get('exintro')).to.equal('true')
                expect(params.get('prop')).to.equal('extracts')
                expect(params.get('format')).to.equal('json')
                expect(params.get('origin')).to.equal('*')
                expect(params.get('titles')).to.exist
            })
        })

        it('should include custom User-Agent header', () => {
            cy.wait('@wikiApi').then((interception) => {
                expect(interception.request.headers)
                    .to.have.property('api-user-agent', 'admin@roadsign.pictures')
            })
        })

        it('should extract title from URL pathname correctly', () => {
            cy.wait('@wikiApi').then((interception) => {
                const url = new URL(interception.request.url)
                const title = url.searchParams.get('titles')

                // Should extract the last segment of the pathname
                expect(title).to.be.a('string')
                expect(title.length).to.be.greaterThan(0)
            })
        })
    })

    context('Wikipedia content formatting', () => {
        beforeEach(() => {
            cy.intercept('GET', '**/en.wikipedia.org/w/api.php*', {
                statusCode: 200,
                body: {
                    query: {
                        pages: {
                            '12345': {
                                pageid: 12345,
                                title: 'Test Highway',
                                extract: 'This is a test highway with multiple sentences. It spans several states. The highway was built in 1950.'
                            }
                        }
                    }
                }
            }).as('wikiApi')

            cy.visit('/highway/az195')
        })

        it('should display entire extract text', () => {
            cy.wait('@wikiApi')

            cy.get('#wiki-search')
                .should('contain', 'This is a test highway')
                .should('contain', 'multiple sentences')
                .should('contain', 'built in 1950')
        })

        it('should append Wikipedia source link to content', () => {
            cy.wait('@wikiApi')

            cy.get('#wiki-search').then($el => {
                const text = $el.text()
                expect(text).to.include('Source:')
                expect(text).to.include('Wikipedia')
            })
        })
    })

    context('Browser compatibility', () => {
        it('should not execute if URL.parse is not supported', () => {
            cy.visit('/highway/az195', {
                onBeforeLoad(win) {
                    // Remove URL.parse to simulate older browser
                    delete win.URL.parse
                }
            })

            // Script should not execute, element should remain empty
            cy.get('#wiki-search').should('be.empty')
        })
    })

    context('Multiple page loads', () => {
        it('should fetch Wikipedia content on each page visit', () => {
            cy.intercept('GET', '**/en.wikipedia.org/w/api.php*').as('wikiApi')

            cy.visit('/highway/az195')
            cy.wait('@wikiApi')
            cy.get('#wiki-search').should('not.be.empty')

            // Navigate away and back
            cy.visit('/')
            cy.visit('/highway/az195')
            cy.wait('@wikiApi')
            cy.get('#wiki-search').should('not.be.empty')
        })
    })
})
