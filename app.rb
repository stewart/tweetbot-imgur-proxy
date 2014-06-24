require 'sinatra'
require 'httmultiparty'

IMGUR_API = "https://api.imgur.com/3/image"
CLIENT_ID = ENV.fetch "IMGUR_CLIENT_ID"

post '/' do
  content_type :json

  unless params[:media]
    p params
    return ""
  end

  file = File.new(params[:media][:tempfile])

  response = HTTMultiParty.post(IMGUR_API, {
    :headers => { "Authorization" => "Client-ID #{CLIENT_ID}" },
    :body => { :image => file }
  })

  body = response.parsed_response

  unless body["status"] == 200
    p response
    return ""
  end

  { :url => body["data"]["link"] }.to_json
end
