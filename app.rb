require 'sinatra'
require 'httmultiparty'

class Imgur
  ENDPOINT = 'https://api.imgur.com/3/image'
  CLIENT_ID = ENV.fetch 'IMGUR_CLIENT_ID'

  def self.upload(filename)
    resp = HTTMultiParty.post(
      ENDPOINT,
      {
        :headers => {
          'Authorization' => 'Client-ID ' + CLIENT_ID
        },

        :body => {
          :image => File.new(filename)
        }
      }
    )

    handle_response resp
  end

  def self.handle_response(response)
    return '' unless response['status'] == 200
    { :url => response['data']['link'] }
  end
end

post '/' do
  content_type :json
  return '' unless params[:media]
  res = Imgur.upload(params[:media][:tempfile])
  res.to_json
end
