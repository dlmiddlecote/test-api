const request = require('supertest');
const expect = require('chai').expect;
const api = request('http://localhost:1234');

describe("resources", function() {
  it('should return a list of resources', function(done) {
    api.get('/resources')
      .set('Accept', 'application/json')
      .expect(200)
      .end(function(err, res) {
        if (err) return done(err);

        res.body.forEach(function(r) {
          expect(r).to.have.property('id');
          expect(r).to.have.property('forename');
          expect(r).to.have.property('surname');
        });

        done();
      });
  });

  it('should return a list of resources with surname middlecote', function(done) {
    api.get('/resources')
      .query({ surname: 'middlecote' })
      .set('Accept', 'application/json')
      .expect(200)
      .end(function(err, res) {
        if (err) return done(err);

        expect(res.body).to.be.an('array').to.not.be.empty;

        res.body.forEach(function(r) {
          expect(r).to.have.property('surname', 'middlecote');
        });

        done();
      });
  });
});
