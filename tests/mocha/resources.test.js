const request = require('supertest');
const expect = require('chai').expect;
const api = request('http://localhost:1234');

describe("resources", function() {
  it('should return a list of resources', function(done) {
    api.get('/resources')
      .set('Accept', 'application/json')
      .expect(200)
      .end(done);
  });

  it('should return a list of resources with surname middlecote', function(done) {
    api.get('/resources')
      .query({ surname: 'middlecote' })
      .set('Accept', 'application/json')
      .expect(200)
      .end(function(err, res) {
        if (err) return done(err);

        expect(res.body).to.be.an('array').to.not.be.empty;

        res.body.forEach(function(ch) {
          expect(ch).to.have.property('surname', 'middlecote');
        });

        done();
      });
  });
});

